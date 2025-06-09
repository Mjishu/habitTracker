import { Text, View } from "react-native";

export default function Index() {
  return (
    <View
      className="flex-1 bg-primary"
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text className="text-5xl text-black-100">Welcome</Text>
    </View>
  );
}
