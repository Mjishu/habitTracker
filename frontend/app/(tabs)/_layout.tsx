import { ImageBackground, Image } from "react-native";
import { View, Text } from "react-native";
import React from "react";
import { Tabs } from "expo-router";
import { images } from "@/assets/constants/images";
import { icons } from "@/assets/constants/icons";
import type { ImageSourcePropType } from "react-native";

const TabIcon = ({ focused, icon, title }: any) => {
  if (focused) {
    return (
      <>
        <ImageBackground
          source={images.highlight}
          className="flex flex-row w-full flex-1 min-w-[112px] min-h-16 mt-4 justify-center 
              items-center rounded-full overflow-hidden"
        >
          <Image source={icon} className="size-5" tintColor="#151312" />
          <Text className="text-secondary text-base font-semibold ml-2">
            {title}
          </Text>
        </ImageBackground>
      </>
    );
  } else
    return (
      <View className="size-full justify-center items-center m5-4 rounded-full">
        <Image source={icon} tintColor="#A8B5db" />
      </View>
    );
};

const _layout = () => {
  return (
    <Tabs
      screenOptions={{
        tabBarShowLabel: false,
        tabBarItemStyle: {
          width: "100%",
          height: " 100%",
          justifyContent: "center",
          alignItems: "center",
        },
        tabBarStyle: {
          backgroundColor: "#0f0d23",
          borderRadius: 50,
          marginHorizontal: 20,
          marginBottom: 36,
          height: 52,
          position: "absolute",
          overflow: "hidden",
          borderWidth: 1,
          borderColor: "#0f0d23",
        },
      }}
    >
      <Tabs.Screen
        name="index"
        options={{
          title: "Home",
          headerShown: false,
          tabBarIcon: ({ focused }) => (
            <TabIcon focused={focused} icon={icons.home} title="Home" />
          ),
        }}
      />
      <Tabs.Screen
        name="habits"
        options={{
          title: "Habits",
          headerShown: false,
          tabBarIcon: ({ focused }) => (
            <TabIcon focused={focused} icon={icons.save} title="Habits" />
          ),
        }}
      />
      <Tabs.Screen
        name="profile"
        options={{
          title: "Profile",
          headerShown: false,
          tabBarIcon: ({ focused }) => (
            <TabIcon focused={focused} icon={icons.person} title="Profile" />
          ),
        }}
      />
    </Tabs>
  );
};

export default _layout;
