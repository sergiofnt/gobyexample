<!DOCTYPE html>
<html>
<head>
	<meta http-eqiv="content-type" content="text/html;charset=utf-8" />
	<title>Go за Прикладом: Цикли (For)</title>
	<link rel=stylesheet href="site.css" />
	<script async src="https://www.googletagmanager.com/gtag/js?id=UA-75170058-3"></script>
	<script type="text/javascript">
	if (window.location.host == "gobyexample.com.ua") {
		window.dataLayer = window.dataLayer || [];
		function gtag(){dataLayer.push(arguments);}
		gtag('js',      new Date());
		gtag('config', 'UA-75170058-3');
	}
	</script>
</head>
<body>
	<div class="example" id="for">
	<h2><a href="./">Go за Прикладом</a>: Цикли (For)</h2>

	<p class="warning">Work in Progress / Сайт в процесі розробки</p>

	
	<table>
		
		<tr>
			<td class="docs"><p><code>for</code> це єдина конструкція циклу в Go. Нижче наведено
три базові типи для циклів <code>for</code>.</p>
</td>
			<td class="code empty leading">
			
			
			</td>
		</tr>
	
		<tr>
			<td class="docs"></td>
			<td class="code leading">
			<a href="http://play.golang.org/p/WTvRHvAO1_A"><img title="Run code" src="play.png" class="run" /></a>
			<div class="highlight"><pre><span class="kn">package</span> <span class="nx">main</span>
</pre></div>

			</td>
		</tr>
	
		<tr>
			<td class="docs"></td>
			<td class="code leading">
			
			<div class="highlight"><pre><span class="kn">import</span> <span class="s">&quot;fmt&quot;</span>
</pre></div>

			</td>
		</tr>
	
		<tr>
			<td class="docs"></td>
			<td class="code leading">
			
			<div class="highlight"><pre><span class="kd">func</span> <span class="nx">main</span><span class="p">()</span> <span class="p">{</span>
</pre></div>

			</td>
		</tr>
	
		<tr>
			<td class="docs"><p>Найбільш розповсюдженим є тип з єдиною інструкцією.</p>
</td>
			<td class="code leading">
			
			<div class="highlight"><pre>    <span class="nx">i</span> <span class="o">:=</span> <span class="mi">1</span>
    <span class="k">for</span> <span class="nx">i</span> <span class="o">&lt;=</span> <span class="mi">3</span> <span class="p">{</span>
        <span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">i</span><span class="p">)</span>
        <span class="nx">i</span> <span class="p">=</span> <span class="nx">i</span> <span class="o">+</span> <span class="mi">1</span>
    <span class="p">}</span>
</pre></div>

			</td>
		</tr>
	
		<tr>
			<td class="docs"><p>Класичний ініціалізація/умова/післяумова <code>for</code> цикл,
також відомий як С-style <code>for</code>.</p>
</td>
			<td class="code leading">
			
			<div class="highlight"><pre>    <span class="k">for</span> <span class="nx">j</span> <span class="o">:=</span> <span class="mi">7</span><span class="p">;</span> <span class="nx">j</span> <span class="o">&lt;=</span> <span class="mi">9</span><span class="p">;</span> <span class="nx">j</span><span class="o">++</span> <span class="p">{</span>
        <span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">j</span><span class="p">)</span>
    <span class="p">}</span>
</pre></div>

			</td>
		</tr>
	
		<tr>
			<td class="docs"><p><code>for</code> без умови буде відбуватись постійно аж
допоки программа не перерве його, за допомогою
ключових слів <code>break</code> (що перериває цикл) або
<code>return</code> (що повертає значення з функції).</p>
</td>
			<td class="code leading">
			
			<div class="highlight"><pre>    <span class="k">for</span> <span class="p">{</span>
        <span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="s">&quot;loop&quot;</span><span class="p">)</span>
        <span class="k">break</span>
    <span class="p">}</span>
</pre></div>

			</td>
		</tr>
	
		<tr>
			<td class="docs"><p>Існує також ключове слово <code>continue</code> - призначенням
якого є перехіду до наступної ітерації циклу.</p>
</td>
			<td class="code">
			
			<div class="highlight"><pre>    <span class="k">for</span> <span class="nx">n</span> <span class="o">:=</span> <span class="mi">0</span><span class="p">;</span> <span class="nx">n</span> <span class="o">&lt;=</span> <span class="mi">5</span><span class="p">;</span> <span class="nx">n</span><span class="o">++</span> <span class="p">{</span>
        <span class="k">if</span> <span class="nx">n</span><span class="o">%</span><span class="mi">2</span> <span class="o">==</span> <span class="mi">0</span> <span class="p">{</span>
            <span class="k">continue</span>
        <span class="p">}</span>
        <span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">n</span><span class="p">)</span>
    <span class="p">}</span>
<span class="p">}</span>
</pre></div>

			</td>
		</tr>
	
	</table>
	
	<table>
		
		<tr>
			<td class="docs"></td>
			<td class="code leading">
			
			<div class="highlight"><pre><span class="gp">$</span> go run <span class="k">for</span>.go
<span class="go">1</span>
<span class="go">2</span>
<span class="go">3</span>
<span class="go">7</span>
<span class="go">8</span>
<span class="go">9</span>
<span class="go">loop</span>
<span class="go">1</span>
<span class="go">3</span>
<span class="go">5</span>
</pre></div>

			</td>
		</tr>
	
		<tr>
			<td class="docs"><p>Ми ще побачимо інші форми <code>for</code> пізніше
коли ми будемо працювати з твердженями <code>range</code>,
каналами та іншими структурами даних.</p>
</td>
			<td class="code empty">
			
			
			</td>
		</tr>
	
	</table>
	


	<p class="next">
		Наступний приклад: <a href="if-else">Оператори Розгалуження - Якщо/Інакше (If/Else)</a>.
	</p>

<p class="footer">
Автор <a href="https://twitter.com/mmcgrana">@mmcgrana</a> | Переклад <a href="https://twitter.com/butuzov">@butuzov</a> | <a href="https://github.com/butuzov/gobyexample/blob/ukrainian/examples/for">source</a> | <a href="https://github.com/butuzov/gobyexample#license">ліцензія</a>
</p>
</div>
</body>
</html>
