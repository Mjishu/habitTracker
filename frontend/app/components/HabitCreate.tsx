import { View, Text, Button, Pressable, Alert, TextInput } from "react-native";
import React, { useState } from "react";
import { CreateHabit } from "@/services/habitServices";

const HabitCreate = () => {
  const [showCreate, setShowCreate] = useState<boolean>(false);

  const CallForm = () => {
    setShowCreate(!showCreate);
  };

  return (
    <View className="w-full h-full flex flex-column justify-center">
      <Pressable onPress={CallForm}>
        <Text className="text-primary bg-accent text-center">Create</Text>
      </Pressable>
      {showCreate === true && <CreateForm setShowCreate={setShowCreate} />}
    </View>
  );
};

interface CreateFormProps {
  setShowCreate: React.Dispatch<React.SetStateAction<boolean>>;
}

const CreateForm: React.FC<CreateFormProps> = ({ setShowCreate }) => {
  const handleSubmit = (): void => {
    //     CreateHabit();
    setShowCreate(false);
  };
  return (
    <View className="w-40 flex flex-col justify-center items-center">
      <Text className="text-accent">Create Form</Text>

      <Text className="text-accent">Name</Text>
      <TextInput
        className="bg-accent rounded-lg w-full"
        placeholder="Brush Teeth"
      />

      <Pressable className="bg-accent mt-4" onPress={handleSubmit}>
        <Text className="text-black-100 rounded-full h-5 w-30 text-center">
          Submit
        </Text>
      </Pressable>
    </View>
  );
};

export default HabitCreate;
