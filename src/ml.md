

# SGD

提供一种 SGD 方法，该方法 实现随机梯度下降


sophisticated algorithm ≤ simple learning algorithm + good training data.

```
>>> import mnist_loader
>>> training_data, validation_data, test_data = mnist_loader.load_data_wrapper()
>>> import network
>>> net = network.Network([784, 30, 10])
>>> net.SGD(training_data, 30, 10, 3.0, test_data=test_data)
Epoch 0: 9058 / 10000, took 8.89 seconds
Epoch 1: 9197 / 10000, took 8.88 seconds
```



# 反向传播算法

反向传播算法最初是在 20 世纪 70 年代引入的， 但直到 著名的 1986 年 论文 大卫 鲁梅尔哈特 _ 杰弗里 辛顿 ,和 罗纳德 威廉姆斯 。 该论文描述了几个 反向传播比以前快得多的神经网络 学习方法，使得使用神经网络来解决问题成为可能 以前无法解决的问题。 今天， 反向传播算法是神经网络学习的主力 网络。

反向传播算法是神经网络训练的关键步骤，通过该算法可以有效地调整神经网络的参数（权重和偏置），以使其逼近期望的输出。该算法的工作原理可以简单概括如下：

    前向传播：首先，通过神经网络的前向传播，将输入数据从输入层经过隐藏层传播到输出层，得到模型的预测输出。

    计算损失：将模型的预测输出与实际期望输出进行比较，计算损失函数(cost function)（如均方误差或交叉熵(cross-entopy)），用于衡量模型预测的准确程度。

    反向传播：根据损失函数，利用链式法则计算损失函数对神经网络中各层参数（权重和偏置）的梯度。这些梯度表示了损失函数对参数的变化敏感程度。

    参数更新：根据梯度下降算法，通过不断迭代，以最小化损失函数为目标，逐步调整神经网络中的参数，使得模型的预测输出逐渐逼近期望输出。

反向传播算法的关键在于通过计算损失函数对参数的梯度，实现了对神经网络参数的自动调整，从而提高了神经网络的准确性和泛化能力。


> an:
> 增加了负反馈调节


# 过拟合

过拟合是指机器学习模型在训练数据上表现良好，但在新数据上表现不佳的不良行为。这种情况发生时，模型会过于复杂，以至于记住了训练数据中的噪音和随机波动，而无法很好地泛化到新的数据集上。过拟合通常是由于训练数据量太小、包含大量无关信息、训练时间过长或模型复杂度过高等原因引起的。

为了检测过拟合，可以使用K折交叉验证等方法。K折交叉验证将训练集等分成K个子集，然后进行K次迭代训练和验证，最终得到模型的平均性能评估。

为了避免过拟合，可以采取一些方法，比如提前停止训练、增加训练数据、数据增强、特征选择、正则化和集成学习等。这些方法有助于降低模型的复杂度，减少噪音的影响，从而提高模型的泛化能力。


> an:
> 原因：前摇过长。
> 解决方法：尽早引入考试和验证，及时纠错




## 避免过拟合：Regularization 正则化

Regularization是机器学习中用来避免过拟合的一种技术。过拟合指的是模型在训练数据上表现良好，但在新数据上表现不佳的情况。Regularization通过对模型的复杂度进行惩罚，来防止模型学习训练数据中的噪声，从而提高模型的泛化能力。

在Regularization中，有两种常见的方法：Ridge Regression和Lasso Regression。Ridge Regression通过向损失函数中添加一个收缩量来修改残差平方和，从而对模型的灵活性进行惩罚。而Lasso Regression则使用绝对值的惩罚来限制模型的复杂度。

这些Regularization技术有助于减少模型的方差，提高模型的泛化能力，而不会显著增加模型的偏差。通过调整Regularization的参数，可以控制对模型偏差和方差的影响，从而在避免过拟合的同时保持模型的重要特性。

除了Ridge Regression和Lasso Regression，Regularization还可以应用于稀疏性正则化、半监督学习和多任务学习等领域。这些Regularization技术的应用有助于提高模型的稳定性和泛化能力，是机器学习中重要的技术手段之一。

> an: 减少模型的复杂性


# 评分函数(score function)和损失函数(loss function)



概述。我们现在将开发一种更强大的图像分类方法，最终将自然地扩展到整个神经网络和卷积神经网络。该方法将有两个主要组成部分：将原始数据映射到类别分数的分数函数(score function)，以及量化预测分数和真实标签之间的一致性的损失函数(loss function)。然后，我们将其视为一个优化问题，目标是更新score function的参数来最小化loss function




# gpu choice

RTX 2070
https://www.thepaper.cn/newsDetail_forward_22138632


3080和4070Ti
https://timdettmers.com/2018/12/16/deep-learning-hardware-guide/



# Adam


https://github.com/kartik4949/tinygrad/commit/26ce2d93c3b8da2eed3ed0620212628e7dcfc459


 George Hotz | Programming | tinygrad and more neural networks from scratch | Part1
直播： https://www.youtube.com/watch?v=Xtws3-Pk69o 35:51

Adam是一种 **优化算法** ，用于更新神经网络权重。它是一种替代传统随机梯度下降程序的优化算法。Adam算法是由OpenAI的Diederik Kingma和多伦多大学的Jimmy Ba在2015年的ICLR论文中提出的。Adam算法的名称源自自适应矩估计（adaptive moment estimation）。Adam算法的优点包括易于实现、计算效率高、内存需求低、适用于大规模数据和参数、适用于非平稳目标、适用于梯度非常嘈杂或稀疏的问题。Adam算法与传统的随机梯度下降不同，它维护每个网络权重的学习率，并在学习过程中分别进行调整。Adam算法结合了AdaGrad和RMSProp的优点，通过计算梯度的一阶矩和二阶矩的指数移动平均值来计算不同参数的自适应学习率。Adam算法在实践中表现良好，与其他随机优化方法相比效果显著。在深度学习领域，Adam算法被广泛应用，并被推荐作为默认的优化方法。



# Dropout

https://github.com/tinygrad/tinygrad/commit/d901ef6b23f76205ad8d47679b75dd0342ef7626 引入tinygrad (Dec 13, 2020)

https://www.cs.toronto.edu/~rsalakhu/papers/srivastava14a.pdf p5 算法

Bernoulli random 0-1随机

Dropout是一种用于深度学习神经网络的正则化方法。在训练过程中，Dropout会随机丢弃神经网络中的一些节点，以防止过拟合。过拟合是指模型在训练数据上表现良好，但在测试数据上表现不佳的情况。Dropout通过随机丢弃节点，使得每次训练时网络的结构都不同，从而增加了模型的鲁棒性，减少了过拟合的风险。

在实践中，Dropout可以通过在每个隐藏层中随机丢弃一定比例的节点来实现。这样做可以使得网络在训练过程中变得更加嘈杂，迫使每个节点以概率性的方式承担更多或更少的输入责任。这种概念表明，Dropout可以打破网络层之间的相互适应，从而使模型更加健壮。

在使用Dropout时，需要设置一个丢弃概率，用于指定在训练过程中丢弃节点的概率。通常情况下，对于隐藏层的节点，丢弃概率可以选择0.5，而对于输入层的节点，通常会选择接近1的概率。在训练结束后，需要对网络的权重进行重新缩放，以便在进行预测时能够得到正确的结果。

总的来说，Dropout是一种简单而有效的正则化方法，可以帮助减少过拟合，提高深度神经网络的泛化能力。


