import csv
import json
from collections import defaultdict

input_file = 'data/subject_questions.csv'
output_file = 'data/formatted_questions.jsonl'

system_content = "This assistant generates questions for competitive exams."

# Dictionary to store questions grouped by subject
subject_questions = defaultdict(list)

# Read the CSV file and group questions by subject
with open(input_file, mode='r') as file:
    csv_reader = csv.reader(file)
    
    next(csv_reader)  # Skip header
    
    for row in csv_reader:
        question, subject = row[0], row[1]
        subject_questions[subject].append(question)

# Limit to 1000 questions per subject
max_questions_per_subject = 1000

# Write the grouped questions for each number from 1 to 60 to the JSONL file in the desired format
with open(output_file, mode='w') as jsonl_file:
    for subject, questions in subject_questions.items():
        # Limit the number of questions processed for each subject
        limited_questions = questions[:max_questions_per_subject]
        
        # For each subject, generate from 1 to 60 questions at a time
        for num_questions in range(1, 61):
            # Loop through the limited questions, making sure to wrap around if there are fewer than `num_questions`
            for i in range(0, len(limited_questions), num_questions):
                batch = limited_questions[i:i+num_questions]

                # Ensure we only add valid batches
                if len(batch) > 0:
                    for question in batch:
                        json_data = {
                            "prompt": f"Generate a question for the subject: {subject}",
                            "completion": question
                        }
                        jsonl_file.write(json.dumps(json_data) + "\n")

print(f"Data formatted and saved to {output_file}")
