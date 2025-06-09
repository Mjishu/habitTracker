import type { User } from "../types/models";

export function GetUser(id: string): User {
  try {
    const res =  fetch(`/api/users/${id}`);
  }

}
