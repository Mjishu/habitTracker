import type { Habit } from "@/types/models";

export async function UserHabits(id: string): Promise<Habit[] | null> {
  const fetchParams = {
    method: "GET",
    headers: {
      "content-type": "application/json",
    },
  };
  try {
    const res = await fetch(`/api/users/habits/${id}`, fetchParams);
    if (!res.ok) {
      throw new Error(`issue getting user habitst: ${res.statusText}`);
    }
    const data = res.json();
    console.log(data);
    return data;
  } catch (err) {
    console.error(`Error trying to get userHabits: ${err}`);
    return null;
  }
}

export async function CreateHabit(newHabit: Habit): Promise<string | null> {
  const fetchParams = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(newHabit),
  };

  try {
    const res = await fetch("/api/habits", fetchParams);
    if (!res.ok) {
      throw new Error(`issue creating a new habit: ${res.statusText}`);
    }
    const data = await res.json();
    console.log(data);
    return data.id;
  } catch (error) {
    console.error(`error trying to create habit ${error}`);
    return null;
  }
}
