/*
  Warnings:

  - You are about to drop the column `voteId` on the `votes` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[pollId,voterId]` on the table `votes` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `voterId` to the `votes` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "votes" DROP CONSTRAINT "votes_voteId_fkey";

-- DropIndex
DROP INDEX "votes_pollId_voteId_key";

-- AlterTable
ALTER TABLE "votes" DROP COLUMN "voteId",
ADD COLUMN     "voterId" TEXT NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "votes_pollId_voterId_key" ON "votes"("pollId", "voterId");

-- AddForeignKey
ALTER TABLE "votes" ADD CONSTRAINT "votes_voterId_fkey" FOREIGN KEY ("voterId") REFERENCES "voters"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
