export interface ItemsResponse {
  items: {
    left_item: ItemResponse;
    right_item: ItemResponse;
  };
}

export interface ItemResponse {
  id: string;
  name: string;
  image_url: string;
}