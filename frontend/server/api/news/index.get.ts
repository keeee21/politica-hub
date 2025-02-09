import prisma from "@/lib/prisma";

export default defineEventHandler(async (event) => {
  return await prisma.news.findMany({
    orderBy: {
      id: "desc",
    },
  });
});
