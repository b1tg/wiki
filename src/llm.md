

# gpt的推理、训练、fine-tuning分别是什么意思

The terms "inference," "training," and "fine-tuning" are fundamental concepts in the field of machine learning and natural language processing. Here's a detailed explanation of each term:

术语“推理”、“训练”和“微调”是机器学习和自然语言处理领域的基本概念。 以下是每个术语的详细解释：




推理：
	含义： 推理是指使用经过训练的模型根据新的输入数据进行预测或生成输出的过程。 当模型被部署并用于执行回答问题、完成提示或生成文本等任务时，它正在参与推理。
	示例： 使用 GPT-3 等预先训练的语言模型来生成类似人类的文本来响应给定的提示，这就是推理的一个示例。

训练：
	含义： 训练涉及通过向机器学习模型提供大量标记数据并允许其学习数据中的模式和关系来教授机器学习模型的过程。 该模型在训练期间调整其内部参数，以最小化其预测与训练数据中的实际标签之间的差异。
	示例： 在文本段落数据集上训练语言模型，以学习人类语言的统计模式和结构。

微调：
	含义： 微调，也称为**迁移学习**，涉及采用预先训练的模型并在特定数据集或任务上对其进行进一步训练，以使其适应新领域或提高其在特定类型数据上的性能。
	示例： 在包含火星城市信息的数据集上微调 GPT-3 等预训练语言模型，使其能够针对有关火星城市的问题生成准确且相关的答案。

https://devv.ai/en/search?threadId=d669udyldpmo


# LoRA


# quantization (量化）

> context:
> You can run 70B LLAMA on dual 4090s/3090s with quantization.

> context 2:
> No, they can't run it. llama 70 with 4 bit quantization takes ~50 GB VRAM for decent enough context size. You need A100, or 2-3 V100 or 4 3090 which all costs roughly roughly $3-5/h



# GGML 和 GGUF



GGML 是创建用于存储 GPT 模型的文件格式的早期尝试。

GGUF 旨在解决 GGML 的局限性并改善整体用户体验。

> GGUF是一种二进制格式，旨在实现快速加载和保存大语言模型，并易于阅读。原来模型是用pytorch保存, 我有的时候也转换成onnx. 但是,想在LLAMA.CPP上跑, 就要用GGUF格式.

# Use cases

##


# fine-tune

https://github.com/brevdev/notebooks/blob/main/llama2-finetune-own-data.ipynb 一些notebook实践

https://news.ycombinator.com/item?id=37484135 精读




# hardware


是的，3090 是当地人工智能社区的一个模因。 此外，支持也令人惊叹，因为它的架构与 A100 基本相同。

3060 也很受欢迎，它是 3090 的一半。

https://news.ycombinator.com/item?id=38589520


 George Hotz | Programming | Mistral mixtral on a tinybox | AMD P2P multi-GPU mixtral-8x7b-32kseqlen
1:24:11 / 2:37:51

# llama


```
git clone --depth=1 https://github.com/ggerganov/llama.cpp
cd llama.cpp/
make

Q:/Downloads/ai/llama.cpp $ ./main.exe  -m "Q:\Downloads\mistral-7b-instruct-v0.1.Q4_K_M.gguf" -p "golang and zig" -n 400 -e
```
