import type { User } from "../types/models";

export function GetUser(id: string): User {
  fetch(`/users/${id}`);
}
