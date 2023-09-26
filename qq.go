package main

import (
	"fmt"
	"regexp"
)

const rexg = `<div id="info" class="">



    
    
  
    <span>
      <span class="pl"> 作者</span>:
        
            
            <a class="" href="/search/%E7%8E%9B%E4%B8%BD%E4%BA%9A%C2%B7%E7%B1%B3%E6%96%AF">[德]玛丽亚·米斯</a>
    </span><br>

    
    
  
    <span class="pl">出版社:</span>
      <a href="https://book.douban.com/press/2880">上海书店出版社</a>
    <br>

    
    
  
    <span class="pl">出品方:</span>
      <a href="https://book.douban.com/producers/561">也人</a>
    <br>

    
    
  
    <span class="pl">副标题:</span> 国际劳动分工中的女性<br>

    
    
  

    
    
  
    <span>
      <span class="pl"> 译者</span>:
        
            
            <a class="" href="/search/%E6%9D%8E%E6%98%95%E4%B8%80">李昕一</a>
        
           /
            
            <a class="" href="/search/%E5%BE%90%E6%98%8E%E5%BC%BA">徐明强</a>
    </span><br>

    
    
  
    <span class="pl">出版年:</span> 2023-8<br>

    
    
  
    <span class="pl">页数:</span> 431<br>

    
    
  
    <span class="pl">定价:</span> 98.00元<br>

    
    
  
    <span class="pl">装帧:</span> 精装<br>

    
    
  
    <span class="pl">丛书:</span>&nbsp;<a href="https://book.douban.com/series/63530">共域世界史</a><br>

    
    
  
    
      
      <span class="pl">ISBN:</span> 9787545822823<br>
<div class="rating_self clearfix" typeof="v:Rating">
      <strong class="ll rating_num " property="v:average"> 8.8 </strong>
      <span property="v:best" content="10.0"></span>
      <div class="rating_right ">
          <div class="ll bigstar45"></div>
            <div class="rating_sum">
                <span class="">
                    <a href="comments" class="rating_people"><span property="v:votes">67</span>人评价</a>
                </span>
            </div>


      </div>
    </div>
<div class="intro">
    <p>【亮点推荐】</p>    <p>★入选20世纪最重要的100本社会学书籍，被誉为《第二性》以来最具影响力的女性主义作品之一。德国著名社会学家、女权主义者玛丽亚•米斯，在第二次国际女权运动的余波中，写就一部女性主义理论范式革新之作，激励全球几代学者和女权活动家。</p>    <p>★首创“家庭主妇化”概念，让“隐形的”女性劳动重见天日；挑战“生产/再生产”的经典定义，促成女性主义与马克思主义里程碑式的合流；戳破国际劳动分工和性别分工的隐秘共谋，直面资本主义文明的剥削本质——女性劳动是资本积累的“内部殖民地”和根基，没有它，资本无限增长的迷梦便难以为继。</p>    <p>★回望历史：中世纪“猎巫运动”怎样戕害女性身体，以血肉为教会、诸侯和资产阶级新贵献祭？观照现实：印度泛滥的“嫁妆谋杀”和“强奸”现象，如何在当代继承了父权制野蛮、血腥的内核？构想未来：废除剥削、回归生存，身体、生产、生活充分自主的女性主义...</p><p><a href="javascript:void(0)" class="j a_show_full">(展开全部)</a></p></div>
</div>`

const scoreRe = `<strong class="ll rating_num " property="v:average">([^<]+)</strong>`

const autoRe = `<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`
const publRe = `<span class="pl">出版社:</span>[\d\D]*?([^<]+)</a>`
const pagesRe = `<span class="pl">页数:</span>([^<]+)<br>`
const priceRe = `<span class="pl">定价:</span>([^<]+)<br/>`
const introRe = `<div class="intro">[\d\D]*?<p>([^<]+)</p>.*</div>`

func main() {
	re := regexp.MustCompile(pagesRe)
	match := re.FindString(rexg)
	fmt.Println(match)
}
