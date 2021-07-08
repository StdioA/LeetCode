import requests
import os
import shutil
import sys


headers = {}

def fetch_question_name(title):
    payload = {
        "operationName":"questionData",
        "variables": {
            "titleSlug": title,
        },
        "query":"query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n }\n}\n",
    }
    res = requests.post("https://leetcode.com/graphql", headers=headers, json=payload)
    question = res.json()['data']['question']
    return int(question['questionId']) if question is not None else 0


def gen_group(id_):
    g = id_ // 100 * 100
    return f"{g}-{g+99}"


def move(fname):
    seg = os.path.splitext(fname)
    question_name = seg[0]
    question_id = fetch_question_name(question_name)
    if not question_id:
        print("Skip", fname)
        return
    folder = gen_group(question_id)
    if not os.path.exists(folder):
        os.mkdir(folder)
    
    newPath = f"{folder}/{question_id}-{fname}"
    print(f"{fname} -> {newPath}")
    shutil.move(fname, newPath)


if __name__ == '__main__':
    move(sys.argv[1])
