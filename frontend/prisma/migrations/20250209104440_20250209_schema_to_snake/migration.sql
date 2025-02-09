/*
  Warnings:

  - You are about to drop the column `createdAt` on the `news` table. All the data in the column will be lost.
  - You are about to drop the column `publishedAt` on the `news` table. All the data in the column will be lost.
  - You are about to drop the column `updatedAt` on the `news` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "news" DROP COLUMN "createdAt",
DROP COLUMN "publishedAt",
DROP COLUMN "updatedAt",
ADD COLUMN     "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "published_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP;
