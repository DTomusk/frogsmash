function makeFavicon(emoji: string) {
  return (
    "data:image/svg+xml," +
    `<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'>
       <text y='.9em' font-size='90'>${emoji}</text>
     </svg>`
  );
}

export const tenantFavicons: Record<string, string> = {
  "frog": makeFavicon("🐸"),
  "book": makeFavicon("📚"),
};
