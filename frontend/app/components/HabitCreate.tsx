import { View, Text, Button, Pressable, Alert, TextInput } from "react-native";
import React, { useState } from "react";
import { CreateHabit } from "@/services/habitServices";

const HabitCreate = () => {
  const [showCreate, setShowCreate] = useState<boolean>(false);

  const CallForm = () => {
    setShowCreate(!showCreate);
  };

  return (
    <View>
      <Pressable onPress={CallForm}>
        <Text className="text-primary bg-accent text-center">Create</Text>
      </Pressable>
      {showCreate === true && <CreateForm />}
    </View>
  );
};

const CreateForm = () => {
  const handleSubmit = () => {
    //     CreateHabit();
  };
  return (
    <View>
      <Text className="text-accent">Create Form</Text>

      <Text className="text-accent">Name</Text>
      <TextInput
        className="bg-accent rounded-full w-100"
        placeholder="Brush Teeth "
      />

      <Pressable className="bg-accent" onPress={handleSubmit}>
        <Text className="text-black-100 rounded-full h-5 w-30">Submit</Text>
      </Pressable>
    </View>
  );
};

export default HabitCreate;
