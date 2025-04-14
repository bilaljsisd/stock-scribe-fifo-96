
// This file contains TypeScript interfaces for Supabase tables

export interface Product {
  id: string;
  name: string;
  sku: string;
  description: string | null;
  units: string | null;
  current_stock: number;
  average_cost: number;
  created_at?: string;
  updated_at?: string;
}

export interface StockEntry {
  id: string;
  product_id: string;
  quantity: number;
  remaining_quantity: number;
  unit_price: number;
  entry_date: string;
  notes: string | null;
  created_at?: string;
}

export interface StockOutput {
  id: string;
  product_id: string;
  total_quantity: number;
  total_cost: number;
  output_date: string;
  reference_number: string | null;
  notes: string | null;
  created_at?: string;
}

export interface Transaction {
  id: string;
  type: 'entry' | 'output';
  product_id: string;
  quantity: number;
  date: string;
  notes: string | null;
  created_at?: string;
  // For entries
  unit_price?: number;
  // For outputs
  total_cost?: number;
  reference_number?: string | null;
  // Ensure compatibility with Go backend which might use reference_id
  reference_id?: string;
}
