export interface Purchasing {
    Purc_id?: number;
    Tran_id?: number;
    Product_ID: number;
    Tracking?: string;
    Create_date?: string;
    Update_date?: string;
    Delflag?: boolean;
}

export interface Transaction {
    Tran_id?: number;
    User_ID?: number;
    Discount: number;
    Total: number;
    Status_code?: number;
    Create_date?: string;
    Update_date?: string;
    Delflag?: boolean;
}

export interface DiscountCode {
    discount_id?: number;
    seller_id?: number;
    discount_code?: string;
    limit?: number;
    used?: number;
    discount_by_percent?: number;
    discount_by_number?: number;
    minimum_total?: number;
    maximum_discount?: number;
    create_date?: Date;
    update_date?: Date;
    delflag?: boolean;
}

export interface Ordering {
    sub_tran_id: number;
    tran_id: number;
    transaction_id: number;
    tranaction_user_id: number;
    user_user_id: number;
    Name: string;
    address: string;
    shopname?: string;
    purchase_transaction_id: number;
    purchase_product_id: number;
    product_product_id: number;
    product_user_id: number;
    product_name: string;
    discount_code: string;
    tracking: string;
    sub_total: number;
    sub_status_code: number;
    status_code: number;
    status_name: string;
    color:string;
    create_date: Date;
    update_date: Date;
    delflag: boolean
}

export interface HistoryOrdering {
  sub_tran_id: number;
  tran_id: number;
  transaction_id: number;
  seller_id: number;
  user_user_id: number;
  name: string;
  address: string;
  shopname: string;
  purchase_transaction_id: number;
  purchase_product_id: number;
  product_product_id: number;
  product_name: string;
  product_user_id: number;
  discount_code: string | null;
  tracking: string | null;
  sub_total: number;
  sub_status_code: number;
  status_code: number;
  status_name: string;
  color: string;
  create_date: Date; // ISO datetime จาก Go (เช่น "2025-08-18T12:34:56Z")
  update_date: Date; // ISO datetime
  delflag: boolean;
}
