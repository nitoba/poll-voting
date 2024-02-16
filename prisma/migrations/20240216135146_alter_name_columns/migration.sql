/*
  Warnings:

  - You are about to drop the column `pollId` on the `poll_options` table. All the data in the column will be lost.
  - You are about to drop the column `voterId` on the `poll_options` table. All the data in the column will be lost.
  - You are about to drop the column `ownerId` on the `polls` table. All the data in the column will be lost.
  - You are about to drop the column `pollId` on the `votes` table. All the data in the column will be lost.
  - You are about to drop the column `pollOptionId` on the `votes` table. All the data in the column will be lost.
  - You are about to drop the column `voterId` on the `votes` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[poll_id,voter_id]` on the table `votes` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `poll_id` to the `poll_options` table without a default value. This is not possible if the table is not empty.
  - Added the required column `owner_id` to the `polls` table without a default value. This is not possible if the table is not empty.
  - Added the required column `poll_id` to the `votes` table without a default value. This is not possible if the table is not empty.
  - Added the required column `poll_option_id` to the `votes` table without a default value. This is not possible if the table is not empty.
  - Added the required column `voter_id` to the `votes` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "poll_options" DROP CONSTRAINT "poll_options_pollId_fkey";

-- DropForeignKey
ALTER TABLE "poll_options" DROP CONSTRAINT "poll_options_voterId_fkey";

-- DropForeignKey
ALTER TABLE "polls" DROP CONSTRAINT "polls_ownerId_fkey";

-- DropForeignKey
ALTER TABLE "votes" DROP CONSTRAINT "votes_pollId_fkey";

-- DropForeignKey
ALTER TABLE "votes" DROP CONSTRAINT "votes_pollOptionId_fkey";

-- DropForeignKey
ALTER TABLE "votes" DROP CONSTRAINT "votes_voterId_fkey";

-- DropIndex
DROP INDEX "votes_pollId_voterId_key";

-- AlterTable
ALTER TABLE "poll_options" DROP COLUMN "pollId",
DROP COLUMN "voterId",
ADD COLUMN     "poll_id" TEXT NOT NULL,
ADD COLUMN     "voter_id" TEXT;

-- AlterTable
ALTER TABLE "polls" DROP COLUMN "ownerId",
ADD COLUMN     "owner_id" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "votes" DROP COLUMN "pollId",
DROP COLUMN "pollOptionId",
DROP COLUMN "voterId",
ADD COLUMN     "poll_id" TEXT NOT NULL,
ADD COLUMN     "poll_option_id" TEXT NOT NULL,
ADD COLUMN     "voter_id" TEXT NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "votes_poll_id_voter_id_key" ON "votes"("poll_id", "voter_id");

-- AddForeignKey
ALTER TABLE "polls" ADD CONSTRAINT "polls_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "voters"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "poll_options" ADD CONSTRAINT "poll_options_poll_id_fkey" FOREIGN KEY ("poll_id") REFERENCES "polls"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "poll_options" ADD CONSTRAINT "poll_options_voter_id_fkey" FOREIGN KEY ("voter_id") REFERENCES "voters"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "votes" ADD CONSTRAINT "votes_poll_id_fkey" FOREIGN KEY ("poll_id") REFERENCES "polls"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "votes" ADD CONSTRAINT "votes_poll_option_id_fkey" FOREIGN KEY ("poll_option_id") REFERENCES "poll_options"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "votes" ADD CONSTRAINT "votes_voter_id_fkey" FOREIGN KEY ("voter_id") REFERENCES "voters"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
