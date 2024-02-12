/*
  Warnings:

  - You are about to drop the column `ownerId` on the `poll_options` table. All the data in the column will be lost.
  - Added the required column `ownerId` to the `polls` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "poll_options" DROP CONSTRAINT "poll_options_ownerId_fkey";

-- AlterTable
ALTER TABLE "poll_options" DROP COLUMN "ownerId",
ADD COLUMN     "voterId" TEXT;

-- AlterTable
ALTER TABLE "polls" ADD COLUMN     "ownerId" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "polls" ADD CONSTRAINT "polls_ownerId_fkey" FOREIGN KEY ("ownerId") REFERENCES "voters"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "poll_options" ADD CONSTRAINT "poll_options_voterId_fkey" FOREIGN KEY ("voterId") REFERENCES "voters"("id") ON DELETE SET NULL ON UPDATE CASCADE;
