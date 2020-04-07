export interface Roles {
  grand?: boolean
}

export interface TeleUser {
  id: number
  lastPhotoId: string
  photoLink: string

  username: string
  last_name: string
  first_name: string
  lang: string,

  activityList: number[]
}
