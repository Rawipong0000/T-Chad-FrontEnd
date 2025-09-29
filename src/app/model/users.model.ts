export interface Users {
    User_ID: number;
    Name: string;
    Last_name: string;
    Email: string;
    Password?: string;
    shopname?: string;
    phone?: string;
    address?: string;
    subdistrict?: string;
    district?: string;
    province?: string;
    postal_code?: string;
    Create_date?: string;
    Update_date?: string;
    Delflag?: boolean;
}

export interface Province {
    province_id: number;
	name_th: string;
	name_en?: string;
}

export interface District {
    district_id: number;
	province_id: number;
	name_th: string;
	name_en?: string;
}

export interface Province {
    province_id: number;
	name_th: string;
	name_en?: string;
}

export interface Subdistrict {
    subdistrict_id: number;
	district_id: number;
	name_th: string;
	name_en: string;
	lat: string;
	long: string;
	zipcode: string;
}