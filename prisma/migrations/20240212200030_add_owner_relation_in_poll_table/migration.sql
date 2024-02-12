/*
  Warnings:

  - Added the required column `ownerId` to the `poll_options` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "poll_options" ADD COLUMN     "ownerId" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "poll_options" ADD CONSTRAINT "poll_options_ownerId_fkey" FOREIGN KEY ("ownerId") REFERENCES "voters"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
