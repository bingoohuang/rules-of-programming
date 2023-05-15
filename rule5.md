## 规则5：优化的第一课是不要优化

我最喜欢的编程任务就是优化。通常意味着使一些代码系统运行更快，但有时我正在优化内存使用或网络带宽或其他一些资源。

这是我喜欢的任务，因为这很容易测量成功。对于大多数编码工作，什么构成成功是模糊的。像这样的书籍试图努力定义什么是良好的代码或良好的系统，但是什么使一行代码好总是不精确的。

这种情况在优化方面并不适用。在那里，答案更加清晰。如果你试图让某些东西运行得更快，你可以直接衡量你的成功。同样，也可以衡量因增加代码大小或复杂性所付出的代价。不用担心半定义的长期效益，也不用信任几年后阅读你的新代码的人会立即理解并被你作为程序员的情感欣赏席卷而过。只有即时的、实际的结果。

在我的热爱优化方面，我并不孤单。实际上，它是如此的有魅力，以至于它引发了每个程序员都知道的一个编程精髓：

> 过早的优化是万恶之源。

顺便说一下，这并不是完整的引用。原始版本由唐纳德·库斯在1974年写下，更加细腻：

> 我们应该忘记小效率，大约97%的时间都是这样：过早的优化是万恶之源。

需要注意的是上下文。在1974年，编译器比现在不太复杂。库斯所说的“小效率”通常是一些棘手的代码片段，用来让编译器生成你想要的代码，如缓存端点以挤出性能的微小部分。

```go
```

或者使用宏来避免函数调用的成本。

```go
```

幸运的是，现在编译器已经足够智能，能够根据你的源代码生成正确的指令，至少对于编写简单直接的代码而言是如此。你越是尝试变得狡猾，编译器就越难以理解你想要什么，因此，老派技巧和复杂的C ++魔法往往最终生成比直接编写逻辑最简表达式更劣质的代码。

然而，智能编译器并没有拯救我们。程序员会本能地担心资源，无论是时间、存储还是带宽，这可能会导致尝试在问题出现之前解决性能问题。

假设你要从列表中选择一个随机项目，这是许多游戏所做的事情。这些物品并没有均等的可能性——每个物品都有相应的权重机会被选中。

```go
```

这很简单。将所有权重相加，然后选择一个不大于总和的随机数。当你减去一个权重时，其中一个权重将导致总和变为负数，选中该加权值。这种情况发生的机会与权重成比例——完成了！

但是，很容易看到这一点，并决定它可能更快。第二次循环值似乎是不必要的。如果你记住权重的运行总和，你可以分治法找到答案，使第二次循环更快：

```go
```

这是一个菜鸟的错误。实际上，这是一个嵌套的菜鸟错误清单。是的，第二个循环现在是O(log N)，而不是线性的，但当第一个循环仍然是线性的时候，那并不重要。你并没有对整体性能产生影响。

实际上，这还不是问题所在。除非你知道你有很多加权随机选择，否则简单的线性循环将更快。直到你达到200个左右的选择，线性循环才比二分查找更快，至少在我的PC上测试结果是如此。这个数字比你想象的还要大，对吧？在此之前，更简单的逻辑和更好的内存访问模式比算法效率更重要。

但这也不是真正的问题。查找的速度多快并不重要——第二个版本分配了内存，这比你所做的任何其他事情都要慢得多。如果你按照先前的编写方式运行这两个函数，第一个版本要比第二个版本快20倍。20倍！


等等——这还不是真正的问题！真正的问题是，无论chooseRandomValue有多快，都无关紧要，只需进行一点点性能剖析就可以快速了解到这一点。你可能每秒钟调用它数百次，但分析器会告诉你它只代表了你总运行时间的微不足道的一部分。Sucker Punch引擎有一些函数每秒钟被调用数百万次；如果你正在写一个游戏，你也会有这样的函数。在性能方面，这些函数很重要，而chooseRandomValue却不重要。

## 优化的第一课

所以，这就是优化的第一课——不要优化。

让你的代码尽可能简单。不要担心它会有多快。它将足够快。如果它不够快，那么让它快起来也很容易。最后那一点——简单代码易于快速优化——是第二课。

## 优化的第二课

想象一下，你有一些简单、可靠的代码，你花了普通的关注写它。你的项目部分运行有点慢，所以你对它进行了测试，发现这一小段代码占了你50%的性能。

这是好消息！如果你能调整那一小段代码的性能，那么你可以将整体性能提高一倍。

顺便说一下，这是相当典型的情况。第一次查看从未进行过优化的一些代码的性能时，通常都是好消息。你很清楚需要处理什么。

坏消息是发现没有什么明显的缓慢，但这在没有经过几轮优化的代码中很少见。

这里有一个经验法则——如果你从未对某个代码进行过优化，那么你可以很容易地将它的速度提高5到10倍。这可能听起来很乐观，但事实并非如此。实际上，在未经优化的代码中有很多低垂的果实。

## 对第二课进行测试

让我们测试一下那个经验法则。想象一下，如果我关于chooseRandomValue错了。它被如此频繁地调用，并且有如此多的选择，以至于它实际上占用了你一半的处理器时间。

现在，如果你从第二个实现开始尝试，我的经验法则就很容易证明。只需切换到一个更简单、无分配模型（就像你的第一个实现），它的运行速度就可以快20倍。经验法则得到证明了！

但那太容易了。假设你从第一个实现开始，所以你没有去除内存分配这个简单的解决方案。实际上，这有点不现实——通常当你查看性能时，你发现有人在循环中分配内存，而这很容易修复。但是让我们假设你运气不好，没有什么简单的方式。

这里有一个优化某些东西的五个步骤。我会着重于性能（“处理器时间”），但是同样的步骤对于任何资源都适用。在下面的步骤中，只需替换网络带宽、内存使用、功耗或任何你试图优化的可测量东西即可。

### 步骤1：测量和归属处理器时间

也就是说，测量花费了多少处理器时间，并将其归因于函数、对象或方便之处。在上面的例子中，我必须已经完成了这个步骤，因为我知道chooseRandomValue占据了我一半的处理器时间。

### 步骤2：确保没有错误

经常发现看起来像是性能问题的东西实际上是个错误。在这种情况下，由于chooseRandomValue实际上吸收了一半的循环次数，我会强烈怀疑某个地方存在错误。我将会认真查看是否所有对chooseRandomValue的调用都是合适的。

也许有人错误地设置了一个循环条件，计数器正在循环回到了初始值。而不是几次迭代，它会循环232次，加减少数不等。那就是很多对chooseRandomValue的调用！（是的，我已经修复了这个确切的错误。）

### 步骤3：测量您的数据

在您知道数据样式之前，不要考虑优化。有多少次调用chooseRandomValue？你要在多少个选项中选择？你是否反复从少数有权重的分布中进行选择，还是它不太可预测？列表中有多少个零权重？您从中选择的值的列表中是否有重复值？

大多数优化都利用了数据或您使用它的某些方面。如果您不彻底了解数据的形状，就无法做出关于优化的正确决策。

### 步骤4：计划和原型

如果您的优化完美无缺——如果它让处理器时间完全为零——那么整体性能会是什么样子？在这种情况下，这意味着chooseRandomValue以零时间运行。如果是这样，您会达到性能目标吗？

如果不是这样，那么您的计划并不够好。您需要确定其他可以优化的代码段。在您知道它是一个能够成功的计划的一部分之前，不要开始优化工作。

有时很难预测完美优化后整体性能会是什么样子。代码与其他代码以不可预测的方式互动。也许chooseRandomValue正在将权值载入处理器的数据缓存中，并且其他一些功能也在使用这些值。在最坏的情况下，您将chooseRandomValue的周期推向零，但整体性能并没有改变。核心问题是将权值加载到数据缓存中——您只是将问题转移给新的罪犯。

寻找一个原型优化的机会。在这种情况下，也许您只需让chooseRandomValue每次返回选择列表中的第一个值。这并不正确，但它很可能为您提供了一个完美最优解的性能预期。

### 步骤5：优化和重复

一旦完成前四个步骤，您可以开始考虑优化。您已经有了关于代码中各个部分应该是多么昂贵的想法，这些想法基于涉及的逻辑和访问的内存数量。也许这些代码或内存访问可以简化或跳过。如果没有一些简单的方式可以加快代码速度，请找到您可以在数据中利用的内容。例如，如果传递给chooseRandomValue的大多数权重为零，则可以利用它。如果有重复值，那么可能是您可以处理的内容。

但是不要盲目行动。寻找一步优化计划的方法：“寻找一些看起来很慢的代码并使其更快”的方法是行不通的。您的直觉可能会误导您，让您误判问题的位置，错误地判断数据的外观，并错误地确定正确的修复方法。

完成第5步后，再次测量您的性能。如果您达到了目标，那就太好了！宣布胜利并停止优化。否则，回到步骤1开始重新评估。 有些步骤第二次可能会更快完成，但是在每个步骤上暂停一下，思考一下到目前为止您学到了什么，这是非常值得的。

应用五步优化过程

好的，我准备开始应用这个过程！我将放下文字处理器，启动开发环境。我将从第一个chooseRandomValue实现开始，应用五步优化过程，看看要获得10倍速度提升需要多少努力。

第一个chooseRandomValue实现是关注代码编写的一个良好示例——它针对简单和清晰进行了优化，这始终是优化的起点。如果我的经验规则是正确的，那么我应该能够在不太费力的情况下获得5倍到10倍的速度提升。

我承认我打这段文字时有点紧张。这可能是非常尴尬的事情。

我已经完成了第一步——我知道我在chooseRandomValue函数中花费了一半的周期。

在第2步，我尽最大努力但没有找到任何错误。所有调用者都有合理的理由进行调用，并且他们没有做任何显然错误的事情。

在第3步，我发现了问题——我在许多情况下调用chooseRandomValue，并传递长的权重和值列表。数据看起来相当随机，尽管权重很小。大多数值都小于5，没有一个大于15。有趣的是，有很多调用，但所有调用都来自于少量静态的分布——也就是说，同样数千个权重和值的列表一遍又一遍地被传递。

在第4步，我创建了一个完美性能版本的chooseRandomValue。在这种情况下，我用一个版本替换了原来的版本，该版本从列表中返回一个随机值，同时忽略权重——很难想象比这更简单的东西了。你可以只返回列表中的第一个值，但那样会跳过似乎是不可避免的随机数生成调用，因此返回一个未加权的随机选择似乎是更好的原型。

我现在正在测试它……代码运行的速度大约是我基准实现的50倍。看起来我预测的5倍到10倍的速度提升还有空间。现在进入第5步——让代码运行得更快！

当你需要让代码运行得更快时，你的第一个冲动可能是实际上让代码运行得更快。做同样的事情，只是更快地做：展开一个循环，使用多媒体指令一次处理多个条目，编写一些汇编语言，将一些数学计算移到循环外部。

这是一种不好的冲动。那些微观优化是你应该尝试的最后一件事，而不是首先尝试的事情。在《铁拳7》中的约两百万行代码中，只有几十个地方我们做了那些微观优化。并不是说我们没有花费大量的精力进行优化——毕竟，我们所做的每一件事都必须在六十分之一秒内完成。我们花费了很多心血让游戏跑得那么快。但是，在极少数情况下，执行同样的操作速度更快并不是我们提高性能的方法。

让代码运行得更快的方法是少做一些操作，而不是将同样的操作速度变快。找出代码正在做但不需要做的事情，或者在可以做一次的情况下做了多次的事情。消除这些代码和内容将使代码运行得更快。

在这种情况下，一个明显的候选者是计算分布的总权重。在chooseRandomValue的第一次实现中，我在每次调用时都要这样做……但是当我在第3步中测量数据时，我发现我正在从有限数量的分布中生成随机值。我可以很容易地为每个分布计算一次总权重，然后在chooseRandomValue中重复使用它：

```go
```

分配内存是昂贵的——这就是为什么第一次尝试优化chooseRandomValue失败的原因。它在每次调用时都分配内存，这完全支配了函数的总成本。然而，在这里，我只是每个分布分配一次内存，而不是每次调用都分配一次。如果我一直在创建新的分布，那么这些分配将是一场灾难，但是我从第3步（那里我测量了数据）知道我只有一个相对较短的分布列表。为这个短列表中的每个分布分配一块内存是可以接受的。

我再次运行代码……它的速度比基线快大约1.7倍。鼓舞人心，但并不是完全胜利。如果您考虑一下这里的数学，您会意识到我最多可能希望获得3倍的加速。以前，平均而言，我在通过权重列表时需要 1.5次，一次完全遍历以计算总重量，然后平均而言需要遍历一半以查找随机值。现在，我只做查找操作。

差别在于内存访问。以前，完全通过权重列表将它们全部拉到某个层次的数据缓存中，因此第二个查找遍历可以快速访问它们。现在第二次遍历需要更多时间来检索值，因此我只获得了1.7倍的加速而不是3倍。

显然的下一步是，现在内存分配是可行的，二分查找更有意义。这不难做到，只是有点微妙：

```go
```

经测试，这个尝试的速度比基线快大约12倍。经验法则验证了！在此处，作者可以想象出一声宽慰的叹息，因为我回到了文字处理器，得到了证明。

大多数情况下，12倍的加速已经足够了。一旦您摘下了最低的果实，就可以继续做其他事情。抵制保持优化的诱惑。很容易沉迷于实际成功的喜悦，追求更多您不需要的性能优势。这个函数不再是性能问题。此时它不再有任何不同于项目中的任何其他函数。它不需要更多的优化。

看，我正在面对这种诱惑。我有更多的想法，可以让chooseRandomValue函数更快。我很好奇到底哪些想法实际上会奏效，而我正在与满足这种好奇心的冲动做斗争。但是，一旦您达到了性能目标，正确的做法是将优化想法作为注释添加到代码中，然后把它放在一边。宣布胜利并继续前进。

我还没有提到一个显而易见的问题。

确实存在一个明显的问题，我还未解答。优化的第一个教训是“不要优化”，对吗？遵循普通的操作规范，编写简单明了的代码，并且相信如果您需要将代码变快5至10倍，那么很容易做到。

但是，如果5至10倍的加速不足够呢？如果您在系统的初始设计中犯了一个巨大的错误，这个错误足够大，使您需要将代码速度变快100倍或1000倍呢？

## 优化没有第三个教训

您可能会认为有第三个优化教训：“不要做傻事。” 如果您要构建一个微秒级别非常重要的高频交易应用程序，那么就不要使用Python构建它。如果您正在定义一些结果结构，并且将在您的C ++代码中随处传递，请不要设计每个副本都进行内存分配。

老实说，我认为第三个教训不存在。程序员总是太担心性能问题，没有必要。

我懂了。我也有同样的弱点。我会为了性能而将复杂性加入代码中，而没有任何性能问题的证据。我一直在做这件事。

也许第三个教训是“不要担心犯错误，因为您不可能犯错而无法修复。”

如果您确实使用Python编写了高频交易应用程序，然后遇到了麻烦，仍有希望。将需要快速处理的部分转换成C ++，将可以缓慢处理的部分保留在Python中。将Python转换为C ++将使代码加快10倍（另一条经验法则），根据我们在这个规则中的实验，一旦转换到C ++，我们可以轻松获得5倍到10倍的速度提升。完成，速度提升50倍至100倍。

实际上，在Sucker Punch，这是我们经常使用的升级路径 - 在我们可爱但相对较慢的脚本语言中编写某些东西的第一个版本，然后如果它成为瓶颈，则将其转换为C ++。我们享受快速尝试想法的好处，知道如果有必要，就有一个更好的性能逃逸路径。

记住，如果您真的犯了一个如此糟糕的错误，以至于您需要找到100倍的性能提升，那么您将很早就知道这一点。这样糟糕的错误不会在草丛中潜伏。它们从一开始就是显而易见的，因此在发现它们之前，您不会深陷其中。所以，再次，不要担心它们。

相信优化的两个教训。编写简单明了的代码，并相信您会找到任何遇到的性能问题的解决方案。

## 插曲：批评上一章

我支持上一章的信息 - 优化的第一个教训确实是“不要优化”。然而！这个强烈的观点，在这本书中众多强烈的观点中单独存在，立即引起了Sucker Punch许多队友的反对。

公平地听取他们充分论证的反对意见是必要的！现在，我将通过与许多持不同意见者之间的假设的苏格拉底式对话来呈现形式，以代表他们的反对观点，为戏剧性目的将他们合并成一个角色。他们都有机会检查本章以确保他们的观点被公正地代表。

反对者：我正式对本章的前提提出异议。

Chris：我觉得这一章只是常识吧。你没看到Knuth的那句名言吗？“我们应该忘掉小效率，也就是说，大约有97%的时间：过早的优化是万恶之源！”

反对者：那句话已经被用来为所有性能极差的代码辩护，而你只是在鼓励更多这样的代码。

Chris：哇。这个反馈中有一种情感的潜流。也许这是因为你不得不花太多时间重新修整其他人的代码，来解决本来不应该存在的性能问题？还有等待受欢迎的视频游戏启动的时间？

反对者：是的。还有。

Chris：而你负责我们代码库中关键性能部分的工作，这可能导致与负责我们用户界面逻辑的人不同的优先级。

反对者：这是真的，不过我要指出，我们都知道有一个游戏的用户界面架构非常糟糕，以至于其性能问题被认为无法修复。整个用户界面必须被抛弃重建，游戏因此错过了发行日期六个月。

Chris：是的。规则20，“做数学”，适用于这种情况。从回顾来看，他们应该意识到他们的架构有多糟糕，并在项目早期进行修复。确实，真正的大性能问题往往会立即出现 - 但只有在你测量它们时。我可以想象第四个优化规则是：“假设你的代码足够快，但仍要进行测量。

反对者：我会稍微感到安心一些。我们能够应对项目末期优化方面的挑战，最大的原因是我们有准确的性能分析工具，并将它们作为日常工程循环的一部分使用。

Chris：是的。我认为这相当于大多数编码团队的测试重心。我们没有太多的单元测试，因为我们愿意让一些错误 sneak in，但我们不愿意被性能问题所困扰。

反对者：尽管如此，我认为你对规则5的论据忽略了一个重要的点。人们很容易将其解读为“不要担心性能优化”，但你实际上是在说，“编写简单代码易于优化。”

Chris：是的，没错。这符合规则1和本书的总体主题：使你的代码尽可能简单，但不要太简单。这种方法的好处之一是，你的代码将易于优化。

反对者：但即使这样，当你编写简单代码时，你也会考虑如果必要的话如何让它更快。当我审查你的代码时，这一点绝对会出现。或者当你审查我的代码时也是如此。实际上，可能两者都是。

Chris：完全正确，我们都会考虑到那个。这不是代码的首要任务 - 正确性和简洁性才是 - 但为优化寻找逃生路径是一个很好的做法，即使它没有被证明是必要的。而且通常并不是必要的。

反对者：确实，性能优化通常不是免费的。如果优化使代码更加复杂，使用更多内存或添加一些预处理步骤，那么性能回报就必须值得。更快的代码并不严格意义上是更好的代码。在这一点上，我们是同意的。

Chris：没错！

反对者：我还想说，尽管简单代码可能易于优化，但慢的代码并不一定简单。实际上，使代码过于复杂是让它变慢的最简单的方法之一。

Chris: 绝对正确。

反对者: 我必须说，这条规则并不能真正体现我所做的大多数优化工作是什么样子的。通常情况下，我不是优化一些新代码 - 我试图从已经被优化过的代码中挤出更多的性能。这很困难。

Chris: 是的。这章确实是关于编写新代码的。

反对者: 是的，但即使是这样，当我向已经知道性能关键的某个系统添加新代码时，我必须从一开始就考虑性能。我不能只写简单的代码然后希望一切顺利。

Chris: 这可能是正确的，或者至少大多数情况下是正确的，这是一个合理的第一步。你同意从一开始就担心性能问题可能导致一些不必要的超优化代码吗？

反对者：虽然我不情愿地接受这一点，但我认为这是不常见的。总的来说，通过不编写立即需要优化的代码，我仍然可以节省时间。

Chris: 我同意。甚至 Knuth 的规则也止步于 97% 对吧？如果你基于过去的经验有信心你正在处理的是 3%，那么在你的第一个实现中考虑性能是合理的。只是在衡量了你的代码并发现问题之前，不要过于迷恋性能。如果你的团队中所有人都认为自己在处理 3%，那么你们都需要更好地对代码进行剖析。

反对者: 在优化过的代码上工作的另一个问题是增益更小。我同意你可以通常将新代码提速 5 或者 10 倍，不需要进行大量的工作。但是在某些时候，你会用完简单的想法，随后性能提升就变得更加困难了。

Chris: 是的，在那个时候规则会改变。你更有可能通过五个小的改变来缩短你的执行时间，而不是通过一个大的改变。但即使是在那个时候，你也必须警觉可能存在更大的算法修复。例如，在第一部《飞天大盗》游戏中，我们花费了数周的时间来优化主绘图循环。我们一次一点点地获得微小的性能，但只是发现切换到一个空间分割系统，可以使其性能提高五倍。

反对者: 那是在我加入之前的事情。不过还是个酷的故事。

Chris: 你觉得五步优化过程怎么样？

反对者: 相当不错。我对这一章的这部分感到满意。

Chris: 我简直无法相信你们没有评论“步骤 2：确保没有漏洞”的杰出洞见。我为这一步感到自豪。

反对者: 我通过不批评它表达了我的赞赏。别期望太多赞美，Chris。我们中没有人想要面对一个更加自信的你。

Chris: 公平的。
