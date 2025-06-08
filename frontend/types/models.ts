export type User = {
  id: string;
  username: string;
  email: string;
  hash_password: string;
  image_src: string;
  max_hp: number;
  current_hp: number;
  base_attack_damage: number;
  xp: number;
  gold: number;
  created_at: string;
  updated_at: string;
};

export type Quest = {
  id: string;
  name: string;
  description: string;
};

export type QuestReward = {
  id: string;
  questId: string;
  itemId: string;
  amount: number;
};

export type QuestLog = {
  id: string;
  questId: string;
  userId: string;
  createdAt: string;
  completedAt: string | null;
};

export type Item = {
  id: number;
  name: string;
  type: string;
  value: number;
  description: string;
  image_src: string;
  created_at: string;
  updated_at: string;
};

export type ItemEffect = {
  id: number;
  itemId: number;
  effect_type: string;
  effect_value: string;
};

export type UserItem = {
  id: string;
  userId: string;
  itemId: number;
  quantity: number;
  acquiredAt: string;
};

export type Enemy = {
  id: string;
  name: string;
  max_hp: number;
  attack: number;
  defense: number;
  xp_reward: number;
  gold_reward: number;
  description: string;
  type: string;
};

export type EnemyReward = {
  id: number;
  enemyId: string;
  itemId: number;
  amount: number;
};

export type UserEnemyEncounter = {
  id: string;
  userId: string;
  enemyId: string;
  current_hp: number;
  encounterStartedAt: string;
  encounterStatus: string;
};

export type Habit = {
  id: string;
  userId: string;
  name: string;
  description: string;
  habitType: string;
  difficulty: string;
  created_at: string;
  updated_at: string;
  effectsStat: string;
};

export type HabitLog = {
  id: string;
  habitId: string;
  userId: string;
  performedAt: string;
  effect: string;
  notes: string;
  mood: string;
};

export type CombatLog = {
  id: string;
  enemyId: string;
  userId: string;
  action: string;
  damageDealt: number | null;
  damageTaken: number | null;
  hpAfterAction: number | null;
  createdAt: string;
};

export type Achievement = {
  id: string;
  name: string;
  description: string;
  rewardType: string;
  rewardValue: number | null;
};

export type UserAchievement = {
  id: string;
  userId: string;
  achievementId: string;
  achievedAt: string;
};
