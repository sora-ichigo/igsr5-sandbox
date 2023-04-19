import os
import openai

openai.api_key = os.getenv("OPENAI_API_KEY")

print("--------------------------")
print("Let's start talking!!")
print("--------------------------")

thread = []
while True:
    try:
        msg = input("You: ")
        thread.append({
            "role": "user",
            "content": msg
        })

        completion = openai.ChatCompletion.create(
            model="gpt-3.5-turbo",
            messages=thread
        )
        
        answer = completion.choices[0].message.content
        thread.append({
            "role": "assistant",
            "content": answer
        })
        print("ChatGPT: {}".format(answer))
    except (EOFError, KeyboardInterrupt):
        # ctrl-cで終了できるように
        break
    

# $ python chat.py
# --------------------------
# Let's start talking!!
# --------------------------
# You: a
# ChatGPT: Sorry, your message is incomplete. Please provide more information or context.
# You: ごめん
# ChatGPT: 大丈夫です。何かご質問やお問い合わせ等ありましたら、お気軽にお知らせください。
# You: さっきなんで謝ったの
# ChatGPT: 私はAIの言語モデルであり、自動翻訳機能を使用しています。あなたの最初のメッセージが不完全な文であったため、私は文脈が不明確であると判断し、追加の情報を求めました。それによってあなたが混乱することを避けるために、私は謝りました。それでも何かご質問があれば、遠慮なくご連絡ください。
# You: マジごめん
# ChatGPT: 全く問題ありません。何か質問があれば、私ができる限りお答えしますのでお気軽にお尋ねください。
# You: ありがとう
# ChatGPT: どういたしまして。何かお力になれることがありましたら、いつでもお知らせください。
# You: 