export interface IConversation {
  id: number;
  name: string;
  description: string | null; 
  is_group: boolean;
  is_user: boolean;
  owner_id: number;
  members: User[]; 
  user_ids: number[];
  created_at: string; 
  updated_at: string; 
  last_message: Message | null;
}

interface Message {
  id: number;
  content: string;
  sender: User;
  created_at: string;
}

interface User {
  id: string
  name: string
  email: string
  avatarURL: string
}
