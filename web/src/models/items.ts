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