export interface IPagination {
  page: number
  size: number
  total: number
  totalPages: number
  hasNext: boolean
  hasPrev: boolean
}

export interface IResponse<T> {
  data: T
  message: string
  status: number
}

export interface IQuery {
  page?: number
  limit?: number
  search?: string
  orderBy?: string
  orderDesc?: boolean
}