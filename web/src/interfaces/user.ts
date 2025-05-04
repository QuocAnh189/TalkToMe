export interface IUser {
  id: string
  name: string
  email: string
  avatar_url: string
  role?: string
  createdAt?: string
  updatedAt?: string
}

export interface IAuth {
  accessToken: string
  refreshToken: string
  user: IUser
}

export interface SignInRequest {
  email: string
  password: string
}

export interface SignUpRequest {
  name: string
  email: string
  password: string
  role: string
  avatar?: any
}

export interface ISearchUser {
  id: string
  name: string
  email: string
  avatar_url: string
  isFriend: boolean
}