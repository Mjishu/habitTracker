import { Text, View } from "react-native";
import HabitCreate from "../components/HabitCreate";

export default function Habits() {
  return (
    <View
      className="flex-1 bg-primary"
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <HabitCreate />
    </View>
  );
}
