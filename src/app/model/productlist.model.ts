export interface Productlist {
  product_id: number;
  product_name: string;
  product_user_id: number;
  user_user_id: number;
  name : string;
  shopname: string;
  price: number;
  description?: string;
  size: string;
  img?: string;
  selling: boolean;
  create_date: string;
  update_date: string;
  delflag: boolean;
}

export interface CartItem {
  id: number;
  name: string;
  price: number;
  size: string;
  seller: string;
  img?: string;
}

export interface GetCartItem {
  product_id: number;
  product_name: string;
  product_user_id: number;
  user_user_id: number;
  name : string;
  shopname: string;
  price: number;
  size: string;
  img?: string;
  selling: boolean;
  delflag: boolean;
}



