// JavaScript Document
$(document).ready(function(){
	function RHeight(){
		/*var docHeight=$(document).height();
		alert(docHeight);
		$('div.right > div.content').css('min-height',(docHeight-40)+'px');*/
		}
	//捕捉一级菜单下拉
	$('ul#ula > li > a').not($('ul#ula li a:first')).click(function(){
			//alert($(this).html());
			var next=$(this).next(); //识别A后面的元素 也就是#ulb
			if(next.attr('id')=='ulb'){ //如果子元素存在
			
			//alert($('ul#ula > li > a').length-1);
			for(var k=0;k<$('ul#ula > li > a').length-1;k++){
				$('ul#ula > li > a').eq(k+1).next().slideUp('fast');
				$('ul#ula > li > a').eq(k+1).removeClass('hover');
					}
			
				if(next.css('display')=='none'){
					$(this).addClass('hover');
					next.slideDown('fast');
					}else{
					$(this).removeClass('hover');
					next.slideUp('fast');
						}	
						
		
				}
					})									
	//自动读取是否下拉并选中
	function selectnav(){
		if(navSelectNo==0){
			return false;
		}else{
		var navli = $('ul#ulb > li');  //读取全部子菜单
		//ja var = 'navSelectNo' is select number
		var navlia = navli.eq(navSelectNo-1).children().eq(0); //选中的A元素
		navlia.addClass('hoverchild'); //选中样式class=hoverchild
		var parul=navlia.parent().parent().parent(); //parent li
		parul.children().eq(0).addClass('hover'); //追加hover
		parul.children().eq(1).show(); //show()
		}
		
		
		
		
		
		
	}
	selectnav(); //AUTO ONCE
	
	//分类目录postcate modify
	$('a.postcatemodify').click(function(){
		$('form#postcatemodify').fadeIn();
		var cid=$(this).attr('cid');
		var cname=$(this).attr('cname');
		$('input#namemodify').val(cname);
		$('input#cidmodify').val(cid);
	})
	
	$('input#postcateModifySubmit').click(function(){
		var cname=$('input#namemodify').val();
		var res=window.confirm('您是否修改：'+cname);
		if(res==true){
			$('form#postcatemodify').submit();
		}else{
			return;
		}
	})
	//menuli modify
	$('a.menulimodify').click(function(){
		$('form#menulimodify').fadeIn();
		var id=$(this).attr('id');
		var name=$(this).attr('name');
		var url=$(this).attr('url');
		$('input#namemodify').val(name);
		$('input#urlmodify').val(url);
		$('input#idmodify').val(id);
	})
	
	$('input#menuliModifySubmit').click(function(){
		var name=$('input#namemodify').val();
		var url=$('input#urlmodify').val();
		var res=window.confirm('您是否修改：'+name+'('+url+')');
		if(res==true){
			$('form#menulimodify').submit();
		}else{
			return;
		}
	})
	//添加菜单menuadd
	$('form#menuadd').submit(function(e){
		var name=$('input#name').val();
		if(name==''){
			alert('菜单名称不得为空');
		}else{
			$(this).submit();
		}
		return e.preventDefault();
	})
	//添加文章
	$('form#postadd').submit(function(e){
		var title=$('input#title').val();
		if(title==''){
			alert('文章标题不得为空');
		}else{
			$(this).submit();
		}
		return e.preventDefault();
	})
	//postsmodify button
	$('input#thumbnailmodify').click(function(){
		$(this).replaceWith('<input type="file" name="thumbnail" />');
	})
	//user add
	$('form#useradd').submit(function(e){
		var phone=$('input#phone').val();
		var password=$('input#password').val();
		if(phone=='' || password==''){
			alert('手机和密码不得为空');
			return e.preventDefault();
		}else{
			$(this).submit();
		}
		
	})
	//安全退出按钮
	// $('.head .pubw .b .bb').click(function(){
	// 		var off=window.confirm('您确定要安全退出吗？');
	// 		if(off){
	// 			alert('您已经退出谢谢');
	// 			}
	// 				})
						   
})