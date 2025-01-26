<script setup lang="ts">

</script>

<template>
  <div class="p-5 bg-gray-50 min-h-screen">
    <h1 class="text-3xl font-bold text-center text-gray-800 mb-8">
      給与明細の控除項目を雑に解説
    </h1>

    <div class="max-w-4xl mx-auto bg-white shadow-md rounded-lg p-6">
      <p class="text-lg text-gray-700 mb-6">
        給与明細に記載されている「控除項目」は、法律で義務付けられた重要な項目です。それぞれの控除が何のために存在するのか、実生活でどのように活用されるのかを具体例を交えながら解説します。
      </p>

      <!-- 健康保険料 -->
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-gray-800 border-b-2 border-gray-200 pb-2 mb-4">
          健康保険料
        </h2>
        <p class="text-gray-700 mb-4">
          健康保険料は、病気やけがで病院にかかったときの医療費を補助するための保険料です。
        </p>
        <ul class="list-disc list-inside text-gray-700 mb-4">
          <li>医療費の自己負担が1~3割で済む（通常は10割）。</li>
          <li>出産時に「出産育児一時金」が支給される。</li>
          <li>病気やケガでやむを得ず休職する際の「傷病手当金」が受け取れる。</li>
          <li>高額な医療費がかかった場合に「高額療養費制度」が適用される。</li>
        </ul>
        <p class="text-sm text-gray-500">
          <em>例:</em> 手術に100万円かかった場合でも、自己負担額はおおよそ30万円程度で済みます。
        </p>

        <!-- 参考にした記事のURL -->
        <h3 class="mt-4">▼参考にした記事</h3>
        <p class="text-sm text-gray-500">
          <a href="https://www.kyoukaikenpo.or.jp/~/media/Files/honbu/event/kohoshizai/01_leaf_iroha23.pdf" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            知っておきたい健康保険のはなし（全国健康保険協会）
          </a>
        </p>
        <p class="text-sm text-gray-500">
          <a href="https://hcm-jinjer.com/blog/kintai/health-insurance-premium_easy-to-understand" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            健康保険料とは何かをわかりやすく解説！国民健康保険料との違いや計算方法を紹介
          </a>
        </p>
      </div>

      <!-- 厚生年金保険 -->
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-gray-800 border-b-2 border-gray-200 pb-2 mb-4">
          厚生年金保険
        </h2>
        <p class="text-gray-700 mb-4">
          厚生年金保険は、将来の年金として老後の生活を支えるために積み立てられる保険料です。厚生年金保険料は毎月の給料やボーナスから計算された金額を事業主と加入者本人で折半し給料やボーナスから天引きされるしくみです。
        </p>
        <ul class="list-disc list-inside text-gray-700 mb-4">
          <li>老後に「老齢厚生年金」として支給される。</li>
          <li>障害を負った場合に「障害厚生年金」が支給される。</li>
          <li>配偶者や子どもがいる場合、「遺族厚生年金」が支給される。</li>
        </ul>
        <p class="text-sm text-gray-500">
          <em>例:</em> 65歳以降に毎月15万円程度の年金を受け取れる場合があります。
        </p>

        <!-- 参考にした記事のURL -->
        <h3 class="mt-4">▼参考にした記事</h3>
        <p class="text-sm text-gray-500">
          <a href="https://www.mhlw.go.jp/nenkinkenshou/document/pdf/mechanism_a4.pdf" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            国民年金と厚生年金の仕組み（厚生労働省）
          </a>
        </p>
        <p class="text-sm text-gray-500">
          <a href="https://kaigo.homes.co.jp/manual/money/pension/kousei/" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            【簡単解説】厚生年金のしくみ｜基礎から保険料、厚生年金基金まで（LIFULL 介護(ライフル介護) ）
          </a>
        </p>
        <p class="text-sm text-gray-500">
          <a href="https://www.freee.co.jp/kb/kb-payroll/how-to-calculate-employees-pension-insurance-premium/" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            「厚生年金とは？」を総まとめ！国民年金との違いや計算方法をわかりやすく解説（freee）
          </a>
        </p>
      </div>

      <!-- 雇用保険 -->
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-gray-800 border-b-2 border-gray-200 pb-2 mb-4">
          雇用保険
        </h2>
        <p class="text-gray-700 mb-4">
          雇用保険は、失業したときや育児・介護休業を取る際に生活を支えるための保険料です。
        </p>
        <ul class="list-disc list-inside text-gray-700 mb-4">
          <li>失業時に「失業手当」が支給される。</li>
          <li>育児休業中に「育児休業給付金」が支給される。</li>
          <li>介護休業中に「介護休業給付金」が支給される。</li>
        </ul>
        <p class="text-sm text-gray-500">
          <em>例:</em> 失業中、月給の約50～80％程度が支給されます。
        </p>

        <!-- 参考にした記事のURL -->
        <h3 class="mt-4">▼参考にした記事</h3>
        <p class="text-sm text-gray-500">
          <a href="https://www.yayoi-kk.co.jp/kyuyo/oyakudachi/koyohoken-01/" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            雇用保険とは？給付の種類や加入条件、保険料の計算方法などを解説（弥生）
          </a>
        </p>
        <p class="text-sm text-gray-500">
          <a href="https://www.jili.or.jp/lifeplan/houseeconomy/1038.html" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            雇用保険について知りたい（公益財団法人 生命保険文化センター）
          </a>
        </p>
      </div>

      <!-- 所得税 -->
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-gray-800 border-b-2 border-gray-200 pb-2 mb-4">
          所得税
        </h2>
        <p class="text-gray-700 mb-4">
          所得税は、国に支払う税金で、収入に応じて課税額が変わります。
        </p>
        <ul class="list-disc list-inside text-gray-700 mb-4">
          <li>年収が高いほど税率が上がる（累進課税）。</li>
          <li>扶養控除や配偶者控除などで減額可能。</li>
          <li>年収が103万円をオーバーすると、超えた分に対して所得税がかかる（「103万の壁」）</li>
        </ul>
        <p class="text-sm text-gray-500">
          <em>例:</em> 年収400万円の場合、約20万円程度が所得税として控除されます。
        </p>

        <!-- 参考にした記事のURL -->
        <h3 class="mt-4">▼参考にした記事</h3>
        <p class="text-sm text-gray-500">
          <a href="https://www.mof.go.jp/tax_information/qanda018.html" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            所得税について教えてください。（財務省）
          </a>
        </p>
        <p class="text-sm text-gray-500">
          <a href="https://www.mmea.biz/2834/" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            所得税とは？わかりやすく解説。いくらから引かれる？（税金・社会保障教育）
          </a>
        </p>
      </div>

      <!-- 住民税 -->
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-gray-800 border-b-2 border-gray-200 pb-2 mb-4">
          住民税
        </h2>
        <p class="text-gray-700 mb-4">
          住民税は、地方自治体に支払う税金で、地域の公共サービスに活用されます。
        </p>
        <ul class="list-disc list-inside text-gray-700 mb-4">
          <li>ゴミ収集や道路整備などの地域サービスの財源。</li>
          <li>所得に応じて課税額が決まる。</li>
          <li>ふるさと納税により控除される</li>
        </ul>
        <p class="text-sm text-gray-500">
          <em>例:</em> 年収300万円の場合、月に約1万円が住民税として控除されます。
        </p>

        <!-- 参考にした記事のURL -->
        <h3 class="mt-4">▼参考にした記事</h3>
        <p class="text-sm text-gray-500">
          <a href="https://www.rakuten-card.co.jp/minna-money/tax_deductible/article_2012_00002/" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            住民税とは何かわかりやすく解説、いつから払わないといけないの？（みんなのマネ活）
          </a>
        </p>
        <p class="text-sm text-gray-500">
          <a href="https://www.nta.go.jp/taxes/shiraberu/taxanswer/shotoku/1155.htm" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            No.1155 ふるさと納税(寄附金控除)（国税庁）
          </a>
        </p>
        <p class="text-sm text-gray-500">
          <a href="https://furunavi.jp/Guide/ResidentTax" target="_blank" rel="noopener noreferrer" class="text-blue-500 underline">
            ふるさと納税で住民税・所得税が控除される仕組みと計算方法（ふるなび）
          </a>
        </p>
      </div>
    </div>
  </div>
</template>
