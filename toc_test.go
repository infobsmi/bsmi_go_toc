package bsmi_go_toc

import (
	"fmt"
	"testing"
)

var exampleHtml = `
<h2 id="mcetoc_1g7bamggu18"><strong>什么是复式记账？</strong></h2>
<p>在其基础上，复式记账是一个看似简单的公式&mdash;&mdash;资产 = 负债 + 权益。用英语&mdash;&mdash;我的意思是，那不是西班牙语或其他任何东西，而是简单的英语&mdash;&mdash;这意味着企业的资产都归某人所有。您要么借钱购买它们，产生更多负债，要么完全拥有它们，产生股权。</p>
<div>&nbsp;</div>
<p>&nbsp;</p>
<p>想象一下，你用自己的 500,000 美元创业。这个等式看起来像 500,000 美元 = 0 美元 + 500,000 美元。如果您借出 100,000 美元的贷款，它将变为 600,000 美元 = 100,000 美元 + 500,000 美元。支付员工 5,000 美元，您最终得到 595,000 美元 = 100,000 美元 + 495,000 美元。</p>
<p>这个高级方程式是复式记账系统使用的所有帐户的汇总。公司制定会计科目表作为建立簿记系统的第一步。这些账户涵盖业务的各个方面，分为资产账户和负债账户。</p>
<h2 id="mcetoc_1g7bamggu19"><strong>借方和贷方&mdash;&mdash;棘手的部分</strong></h2>
<p>双重输入很容易且相对直观，但有一个明显的例外。在我们的日常工作和演讲中，财务&ldquo;借&rdquo;是取款，&ldquo;贷&rdquo;是加。在复式簿记中，情况并非如此。</p>
<h3>复式记账案例</h3>
<p>在该系统中，&ldquo;借方&rdquo;一词仅表示在两列输入系统的左列中进行输入，而&ldquo;贷方&rdquo;表示在右侧输入。如果你能把它卡在你的脑海里，那一切都是有道理的。</p>

<h4>复式记账案例 jj</h4>
<p>对于资产和费用，借方的条目表明账户余额增加。对于负债、权益和收入，增加记录在贷方栏中。如果系统的这一部分是压倒性的，只需制作一个备忘单。</p>

<h4>复式记账案例 kk</h4>
<h3>复式记账案例 kk</h3>
<h4>复式记账案例 kk</h4>
<h3>复式记账案例 kk</h3>
<h4>复式记账案例 kk</h4>
<h4>复式记账案例 kk</h4>
<h3>复式记账案例 kk</h3>
<h4>复式记账案例 kk</h4>
<h3>复式记账案例 kk</h3>
<h4>复式记账案例 kk</h4>
<h5>复式记账案例 kk</h5>
<h3>复式记账案例 kk</h3>
<h4>复式记账案例 kk</h4>
<h3>复式记账案例 kk</h3>
<h4>复式记账案例 kk</h4>
<p>因此，例如，当您的公司购买拖拉机时&mdash;&mdash;为什么你的甜甜圈企业需要拖拉机我无法理解，但这是一个会计部分，而不是战略部分&mdash;&mdash;你将增加一种资产并减少另一种资产。具体来说，您将增加&ldquo;机械&rdquo;等资产类型并减少&ldquo;现金&rdquo;。</p>
<p>在此示例中，您将向机械账户进行借记分录&mdash;&mdash;增加的资产获得借记分录&mdash;&mdash;以及对现金账户的贷记分录&mdash;&mdash;减少的资产获得贷记分录。</p>
<p>回顾会计等式，您的数字实际上不会改变，因为您只增加和减少了资产。结果，增加和减少都发生在等式的同一侧。</p>
<p>如果您使用赊购方式进行购买，您将保留机械账户的借方分录，但您将拥有代表贷款的负债的贷方分录。</p>
<h2 id="mcetoc_1g7bamggu1a"><strong>例子</strong></h2>
<p>这里有五个例子可以帮助你明白这一点。</p>
<p>用现金购买一台机器。如上所述，资产增加时输入借方列，减少时输入借方列。因此，您将获得如下所示的余额。</p>
<table style="border-collapse: collapse; border-color: #000000; border-style: solid;" border="1">
<tbody>
<tr>
<td>帐户</td>
<td>借方</td>
<td>贷方</td>
</tr>
<tr>
<td>机械（资产）</td>
<td>5,000 美元</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>现金（资产）</td>
<td>&nbsp;</td>
<td>5,000 美元</td>
</tr>
</tbody>
</table>
<p>以下是支付员工工资的方式 - 详细信息将取决于您的会计科目表。您通过借记分录增加了开支，通过贷记分录减少了现金。</p>
<table style="border-collapse: collapse; border-style: solid;" border="1">
<tbody>
<tr>
<td>帐户</td>
<td>借方</td>
<td>贷方</td>
</tr>
<tr>
<td>工资（费用）</td>
<td>2,200 美元</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>工资税（费用）</td>
<td>300 美元</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>现金（资产）</td>
<td>&nbsp;</td>
<td>2,500 美元</td>
</tr>
</tbody>
</table>
<p>减记存货价值是一件简单的事情，即通过输入贷方并将其与销货成本中的借方或单独的存货冲销账户相匹配来降低您的库存账户的价值。</p>
<table style="border-collapse: collapse; border-style: solid;" border="1">
<tbody>
<tr>
<td>帐户</td>
<td>借方</td>
<td>贷方</td>
</tr>
<tr>
<td>销货成本（费用）</td>
<td>4,500 美元</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>库存（资产）</td>
<td>&nbsp;</td>
<td>4,500 美元</td>
</tr>
</tbody>
</table>
<p>对于向投资者出售股票，您将产生现金并增加股权。</p>
<table style="border-collapse: collapse; border-style: solid;" border="1">
<tbody>
<tr>
<td>帐户</td>
<td>借方</td>
<td>贷方</td>
</tr>
<tr>
<td>现金（资产）</td>
<td>50,000 美元</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>普通股（股权）</td>
<td>&nbsp;</td>
<td>50,000 美元</td>
</tr>
</tbody>
</table>
<p>最后，如果您向银行偿还贷款，您将减少手头的现金，同时也减少贷款的负债。</p>
<table style="border-collapse: collapse; border-style: solid;" border="1">
<tbody>
<tr>
<td>帐户</td>
<td>借方</td>
<td>贷方</td>
</tr>
<tr>
<td>应付贷款（负债）</td>
<td>2,500 美元</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>现金（资产）</td>
<td>&nbsp;</td>
<td>2,500 美元</td>
</tr>
</tbody>
</table>
<h2 id="mcetoc_1g7bamggu1b"><strong>最后的笔记</strong></h2>
<p>对于几乎所有企业来说，所有这些复式工作都将在您的会计软件包的幕后进行。应计会计&mdash;&mdash;复式记账的替代方法&mdash;&mdash;根本不用于现代会计软件。</p>
<p>即便如此，了解复式记账的理论和流程的好处可以帮助您更好地了解您的企业财务运作方式。这是每个小型企业都可以使用的东西。</p>`

func TestTocGen(t *testing.T) {
	oc, ob, _ := TocGenerate(exampleHtml)
	fmt.Println("outStr: " + oc)
	fmt.Println("afterHtml: " + ob)
	fmt.Println("hello world")
}
