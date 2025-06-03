import { useState } from 'react';

type QuestInfo = {
  id: number;
  name: string;
  reward: string[];
};

function findQuest(quest: number): QuestInfo {
  let info = {
    id: 0,
    name: '',
    reward: [],
  };

  try {
    fetch(`newUrl/${quest}`)
      .then((response) => response.json())
      .then((data) => {
        info = data;
      })
      .catch((error) => {
        console.error('Fetch error:', error);
      });
  } catch (error) {
    console.error('Try-catch error:', error);
  }

  return info;
}

export function QuestGenerator(questNumber: number) {
  const [questInfo, setQuestInfo]: QuestInfo = useState({ id: 0, name: '', reward: [] });

  setQuestInfo(findQuest(questNumber));

  return <></>;
}
