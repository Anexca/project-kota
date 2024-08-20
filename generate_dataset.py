import csv
import json

input_file = 'subject_questions.csv'
output_file = 'formatted_questions.jsonl'

system_content = "This assistant generates questions for competitive exams."

with open(input_file, mode='r') as file:
    csv_reader = csv.reader(file)
    
    next(csv_reader)
    
    with open(output_file, mode='w') as jsonl_file:
        for row in csv_reader:
            question, subject = row[0], row[1]
            message = {
                "messages": [
                    {"role": "system", "content": system_content},
                    {"role": "user", "content": f"Generate a question for the subject: {subject}"},
                    {"role": "assistant", "content": question}
                ]
            }

            jsonl_file.write(json.dumps(message) + "\n")

print(f"Data formatted and saved to {output_file}")
