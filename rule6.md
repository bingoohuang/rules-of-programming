# 第六条原则：代码审查有三个好处

在我全职程序员的 30 多年中最大的变化之一是逐渐普及各种形式的代码审查。

直到 90 年代初，我甚至都没有听说过代码审查。我不是说它们没有出现，因为当然有，但除了失败不是一个选择的情况，比如医疗设备固件或火箭控制代码，它们并不普及。你知道的，那些会导致人死亡的东西。

30 年前大多数程序员都觉得让别人查看你的代码是……入侵性的。当然，如果你正在与他人合作，你必须至少看一下队友的代码接口，以确定如何与其交互，并且你可能会单步执行某些人的代码，但实际上逐行查看代码并对其进行评判感觉很奇怪，就像阅读别人的日记，或者（用现代的术语）偶然发现了某人的浏览历史一样。

无论如何，90 年代初我转到了微软的一个有代码审查政策的团队。幸运的是，我的项目在这个团队的大计划中如此微不足道，以至于我们的项目团队被完全忽视了。我们不得不决定我们的代码审查流程。我甚至不确定大团队的实际正式代码审查流程是什么；我们只是按照我们认为有意义的方式进行，没有人对我们进行过审查。我肯定不会请求指导，因为我害怕会有一些可怕的流程强加给我们。比起请求许可，向他人请求原谅更好。

出乎我的意料，我发现代码审查立即和无可否认地有用。自那时以来，我一直与我的团队一起这样做——但不是因为我预期的原因。

代码审查最明显的原因是在项目检入之前检测错误。如果你的代码审查过程在任何方面都是合理的，那么进行审查的人会经过准备，能够理解正在检查的代码。也许他们曾经是这段代码实现的一部分，或者是某个新代码所依赖的其他代码的专家，或者是正在审查的代码的频繁用户。在任何情况下，他们可能会发现问题，比如你忽略或破坏了某些假设，误用了你正在调用的某个代码片段，或者可能会改变系统的行为，从而会破坏审查员正在处理的某个其他代码片段。

这种情况会发生吗？代码审查是否真的会找出错误？当然会，至少根据我在我的团队上如何进行代码审查的经验来说，能找到一些问题。

这是一个重要的警告。你从代码审查中获得的价值将取决于你投入多少时间和精力以及你如何进行代码审查。以下是大多数 Sucker Punch 代码审查目前的快速说明：

- 这种代码审查是实时的——至少在疫情前，两个人坐在同一台计算机旁。	
- 这是非正式的。当你有代码需要进行审查时，你走到一个合适的审查人的办公室，请求他们给你进行审查。我们的社交契约是，当有人请求审查时，你同意，除非有着真正紧急的情况。
- 审查员使用一种 diff 工具浏览变更，审查的人对变更进行评论。这是一个对话，审查员会询问问题，直到他们满意地理解所作的更改，提出建议，确定需要测试哪些内容，并讨论其他方法。让审查人员主导审查通常是一个错误；审查员太容易只接受审查人员所说的，而没有自己思考。
- 进行审查的人需要记录所有建议的更改和额外需要运行的测试。社交契约是，至少默认情况下，所有建议都要被纳入到代码中。
- 根据更改的范围，代码审查可能需要五分钟或五个小时。很少会出现一次代码审查没有进行至少一次或两次的更改然后再进行提交。经过一次大规模的代码审查，可能要编写成页数和页数的备注。
- 通常，一次代码审查就足够了。在进行适当的更改并运行额外的测试后，审查人员进行提交。有时，如果这是一个带有很多审查备注的大规模更改，审查人员可能会重新审查更新后的更改。如果原始的代码审查人员对更改的某些部分不确定，他们可能会建议团队中的另一个人也审查此代码。但通常情况下，代码审查 + 纳入更改 + 提交就足够了。

通过这个过程，我们确实发现了 bug……但同样，并不是你期望的方式。以下是我们发现代码审查中 bug 的三种基本方式，按照发现 bug 的频率大致排序，从最常见到最不常见：

- 在请求审查之前，你先仔细查看 diff，确保在向别人展示之前修正了任何可能令人尴尬的地方。在自我审查过程中，你发现了一个 bug，比如说一个你错过的错误情况。你在别人看到之前解决了问题。	
- 在审查过程中，你正在向审查员讲解代码的特定部分……因为被迫解释你的方法会帮助你理解为什么它有缺陷。你向审查员指出了 bug，讨论随之展开，你记下了注释，然后继续进行。或者，如果你发现的缺陷足够严重，你可以完全放弃代码审查，先进行必要的全面更改，然后重新开始审查。
- 在审查过程中，审查员看到了你错过的问题。或者，你描述你所做的事情的方式使得清楚地意识到你对你所调用的某些代码有误解。你讨论可能存在的问题，同意问题存在，并做出注释。

审查员仅仅通过盯着代码进行深思熟虑就发现了 bug 是非常罕见的。代码审查过程本身往往会将它们展现出来，或者在准备阶段或讨论更改的过程中展现出来。这就是为什么代码审查应该是一种对话的过程——解释事情的过程和理解这种解释，揭示了审查员和审查人之间任何不匹配的假设。这对于发现 bug 是好的，但也对于知道何处需要注释或更改名称是有益的。

我们必须指出我们的代码审查过程的无可争议的局限性。我们所有的代码中的每一个 bug 都成功地遁入了代码审查，而我们有成千上万的 bug！我们不会为代码审查要求做任何例外——我们检查的每一行代码都已经经过了审查，所以每一个 bug 都是在它被检查之前被多个人忽略的。代码审查可以找到 bug，但它们肯定不是所有 bug 中的全部。

代码审查是一种低效的寻找 bug 的方式。但我们仍在进行它们。那是因为发现 bug 只是我们进行代码审查的原因之一，而且这并不是最重要的原因。

## 代码审查是关于分享知识的

这里有更重要的理由来进行代码审查——如果进行得当，它们是在整个团队中传播知识的极好方式。

对于 Sucker Punch 的团队来说，这尤其重要，因为我们对任务的指派非常灵活，程序员在我们的代码库的不同部分之间相当自由地移动。如果每个程序员都有关于代码库的不同部分的基本知识，这将会更好。代码审查是传播这种知识的好方法。

想象一下，如果将你的团队中的程序员随意分成“初级”和“高级”两组，大致根据对代码库的熟悉程度。高级程序员对代码库非常了解，而初级程序员正在学习它的各个方面。我们的代码审查涉及两个人，因此评审者和审查人员之间有四种可能的资历组合。只有三种是有用的，如表 6-1 所示。

表 6-1.  代码审查分类法

|            | 高级评审者 | 初级评审者 |
| ---------- | ---------- | ---------- |
| 高级审阅人 | 有用       | 有用       |
| 初级审阅人 | 有用       | 禁止       |

如果一位高级程序员审查一位初级程序员的工作，他们能够很好地看到问题——不仅是审阅的代码中的 bug，而且是初级程序员存在的一般误解。也许初级程序员没有正确地遵循团队的格式标准，或者他们过早地概括了解决方案，或者他们在一个简单的问题上编写了一个复杂的解决方案。这些并不是 bug，但违反了编程规则，会降低代码的质量，因此高级程序员应在代码审查中注意到这些问题并加以解决。

如果一位初级程序员审查一位高级程序员的工作，他们不太可能发现问题，但他们更可能提出问题以弄清发生了什么。在回答这些问题的过程中，被评审人有助于评审人理解代码的上下文，让他们更好地了解代码库的所有组成部分。评审人能够看到和询问优秀代码的示例，即正确格式化的、适当工程化的、结构清晰且命名清晰的代码。

把这两种初级-高级交互视为团队新成员教育过程的一部分。要有效，新成员需要了解所有组成部分如何拼凑在一起、您的团队如何编写代码以及为什么要这样做。代码审查是向团队新成员转移所有这些非正式知识的绝佳方式。

第三种有用的组合是一位高级程序员审查另一位高级程序员的代码。这是发现 bug 和检查两个程序员如何将更改整体地加入进去的绝佳机会，讨论该领域的未来工作，识别可能要运行的额外测试，并确保至少有两个人理解被检查的代码行。

最后一种组合是一位初级程序员审查另一位初级程序员的工作，这是毫无用处的。事实上，它可能会带来负面影响。当两个程序员都是初级的时候，我刚才讨论的所有优点都不复存在。没有知识传递，没有足够上下文来找到 bug，也没有使用代码审查作为未来方向讨论的跳板。最糟糕的是，两个初级程序员之间会相互反射半成品的意见，直到它们看起来像官方团队政策。当 Sucker Punch 代码中出现奇怪的范式和惯例（尽管我们尽最大努力避免），这通常是两个初级程序员来回反弹审阅的结果。因此，我们禁止这种代码审查方式。

## 代码审查的价值

我们发现 bug，传递知识。这可能足以证明我们在代码审查中投入的努力——通常相当于写代码所用时间的5%至10%。但是，代码审查还有一个更重要的好处，可能是全部好处中最重要的，它完全是社交性的：

> 每个人都会编写更好的代码，如果他们知道有人会看它的话。


在代码审查之前，我们会更好地遵循格式和命名约定。我们不会采取捷径，或者把任务留到以后。我们的注释会更清晰。我们是用正确的方法解决问题，而不是通过 hack 和变通方法。我们会记得删除用于诊断问题的临时代码。

所有这些都会在实际代码审查之前发生，这是因为作为程序员，我们给自己施加压力，以编写我们自豪并且乐意向同行展示的工作。这是一种健康的同伴压力。我们编写更好的代码，随着时间的推移，这将导致更健康的代码库和更高效的团队。

## 代码审查天生就是社交行为

总结一下，进行良好的代码审查有以下三个好处：

  1. 你会发现一些错误。
  2. 每个人都会更好地理解代码。
  3. 人们会编写他们乐于分享的代码。

看，代码审查就像任何流程一样。如果您要花时间进行审查，就要使这个过程产生成果。这意味着需要考虑你从中得到了什么以及为什么要进行审查。摒弃并不起到帮助作用的部分，加倍努力推动有效的方法。这样你可以获得更多的价值，或者花费较少时间以获得相同的价值。

除非你像一些程序员进行的配对编程一样，否则编写和调试代码通常都是一个单独的行为。一个孤独的战士，在键盘前独自战斗，战胜程序缺陷和固执的库。

代码审查并不是一个孤独的过程，它们的大部分价值来自审查者和被审查者之间的社交互动。当你在解释一行代码时发现了一个错误，当你解释了一个代码片段，让审查者在下一次调用时正确使用它，当你在请求审查之前清除不想让任何人看到的 hack，或者当你从被审查者的技术解释中学到了更简单的解决问题的方法时，你就会意识到其中的价值。

知道代码审查的价值来自于社交互动，两个人通过交流实现修改时，你应该确保你的代码审查流程鼓励这种互动。如果审查很安静——如果审查者默默地翻阅差异，并在被审查者默默地观察时偶尔发出呻吟声——那么一定出了问题。是的，这仍然是一个代码审查，但你错失了审查可能提供的真正价值。

如果你的所有代码审查都变成了争论，那么你肯定做错了！一个不接受审查者建议的被审查者将不会学到任何东西；如果审查者不努力理解被审查者编写代码的原因，他也不会学到任何东西。而且，在任何情况下，代码审查都不是关于项目方向、团队惯例或哲学进行争论的场所。要在团队中解决这些问题；在一个两个人之间的争吵中，你不会得到任何解决方案。

一个健康的代码审查加强了你的代码库，同时也加强了你团队的联系。这是一个专业和开放的对话，两个参与者都会从中学到东西。
