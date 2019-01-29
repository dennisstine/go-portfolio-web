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
    submitHandler: function(form) {

      $.ajax({
        url: '/messages',
        type: "post",
        dataType: 'text',
        data: $("#contact-form").serialize(),
        success: function() {
          $('#modal-success').toggleClass('is-active');
        },
        error: function() {
          $('#modal-error').toggleClass('is-active');
        }
      });
    }
  });

  $('#close-success').click(function(e) {
    $('#modal-success').toggleClass('is-active');
  });

  $('#close-error').click(function(e) {
    $('#modal-error').toggleClass('is-active');
  });
});
