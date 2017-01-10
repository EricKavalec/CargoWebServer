// bs-example-navbar-collapse-1
$('#bs-example-navbar-collapse-1 li').click(function(e) {
    $('#bs-example-navbar-collapse-1 li.active').removeClass('active');
    var $this = $(this);
    if (!$this.hasClass('active')) {
        $this.addClass('active');
    }
    e.preventDefault();
    hidePages(this.id)
});

$('#a-cargo').click(function(e) {
    $('#bs-example-navbar-collapse-1 li.active').removeClass('active');
    e.preventDefault();
    hidePages(this.id)
});

function hidePages(exceptId){
	$('.pageRow').addClass('hidden')
	$('#page-' + exceptId.split('-')[1]).removeClass('hidden')
}


main()

function main() {
     
}
