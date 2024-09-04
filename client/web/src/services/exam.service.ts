import axiosInstance from "./base";

export const getQuestions = async () => {
  return {
    data: [
      {
        id: 1,
        raw_exam_data: {
          hints: [
            "Discuss the potential benefits, such as increased efficiency and personalized services.",
            "Analyze the challenges, including job displacement and ethical considerations.",
            "Provide examples of AI applications in banking and their implications.",
          ],
          topic:
            "The impact of artificial intelligence on the future of banking",
          type: "essay",
        },
        user_attempts: 0,
        max_attempts: 2,
        duration_minutes: 30,
        number_of_questions: 1,
        created_at: "2024-09-02T19:27:00.239927Z",
        updated_at: "2024-09-02T19:27:00.239927Z",
      },
    ],
  };
  const response = await axiosInstance.get("/exams/banking/descriptive");
  return response.data;
};
