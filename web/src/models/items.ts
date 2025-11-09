export interface Item {
  id: string;
  name: string;
  image_url: string;
}

export interface ItemsResponse {
  items: {
    left_item: Item;
    right_item: Item;
  };
}

export interface LeaderboardResponse {
  data: LeaderboardItem[];
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

export interface LeaderboardItem {
  rank: number;
  id: string;
  name: string;
  image_url: string;
  score: number;
}