$(function() {

  AOS.init();

  $('.navbar-burger').click(function () {
    $('#main-navbar, .navbar-burger').toggleClass('is-active');
  });

  $('a[href*="#"]').click(function (e) {
    e.preventDefault();

    $('html, body').animate({
      scrollTop: $($(this).attr('href')).offset().top
    }, 500);

    $('#main-navbar, .navbar-burger').toggleClass('is-active');
  });

  setTimeout(function(){
    $('#arrow').fadeIn(500);
    $('#arrow').addClass('arrow');
  }, 1500);

  $('#contact-form').validate({
    errorClass: 'is-danger',
    rules: {
      name: {
        required: true,
      },
      email: {
        required: true,
        email: true
      },
      subject: {
        required: true
      },
      text: {
        required: true
      }
    },
    messages: {
      name: 'Please tell me your name',
      email: 'Your email address to reply to',
      subject: 'Let me know why you\'re contacting me',
      text: 'More details go here'
    },
    submitHandler: function() {

      $.ajax({
        url: '/messages',
        type: "post",
        dataType: 'text',
        data: $("#contact-form").serialize(),
        success: function() {
          $('#modal-success').addClass('is-active');
          setTimeout(function(){
            $('#modal-success').removeClass('is-active');
          }, 3000);
        },
        error: function() {
          $('#modal-error').addClass('is-active');
          setTimeout(function(){
            $('#modal-error').removeClass('is-active');
          }, 3000);
        }
      });
    }
  });

  $('#close-success').click(function(e) {
    $('#modal-success').removeClass('is-active');
  });

  $('#close-error').click(function(e) {
    $('#modal-error').removeClass('is-active');
  });
});
