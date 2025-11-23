export interface LeaderboardResponse {
  data: LeaderboardItemResponse[];
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

export interface LeaderboardItemResponse {
  rank: number;
  id: string;
  name: string;
  image_url: string;
  score: number;
  created_at: string;
  license: string;
}