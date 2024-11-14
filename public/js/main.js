/* ===================================================================
 * Count - Main JS
 *
 * ------------------------------------------------------------------- */

(function($) {

    "use strict";
    
    var cfg = {
        scrollDuration : 800,
        mailChimpURL   : 'https://facebook.us8.list-manage.com/subscribe/post?u=cdb7b577e41181934ed6a6a44&amp;id=e6957d85dc'
    },

    $WIN = $(window);

    var doc = document.documentElement;
    doc.setAttribute('data-useragent', navigator.userAgent);

    var ssPreloader = function() {
        $("html").addClass('ss-preload');

        $WIN.on('load', function() {
            $("#loader").fadeOut("slow", function() {
                $("#preloader").delay(300).fadeOut("slow");
            }); 
            $("html").removeClass('ss-preload');
            $("html").addClass('ss-loaded');
        });
    };

    var ssInfoToggle = function() {
        $('.info-toggle').on('click', function(event) {
            event.preventDefault();
            $('body').toggleClass('info-is-visible');
        });
    };

    var ssSlickSlider = function() {
        $('.home-slider').slick({
            arrows: false,
            dots: false,
            autoplay: true,
            autoplaySpeed: 3000,
            fade: true,
            speed: 3000
        });
    };

    var ssPlaceholder = function() {
        $('input, textarea, select').placeholder();
    };

    var ssFinalCountdown = function() {
        var finalDate =  new Date("March 25, 2021 15:37:25").getTime();
        $('.home-content__clock').countdown(finalDate)
        .on('update.countdown finish.countdown', function(event) {
            var str = '<div class=\"top\"><div class=\"time days\">' +
                      '%D <span>day%!D</span>' + 
                      '</div></div>' +
                      '<div class=\"time hours\">' +
                      '%H <span>H</span></div>' +
                      '<div class=\"time minutes\">' +
                      '%M <span>M</span></div>' +
                      '<div class=\"time seconds\">' +
                      '%S <span>S</span></div>';
            $(this)
            .html(event.strftime(str));
        });
    };

    var ssAjaxChimp = function() {
        $('#mc-form').ajaxChimp({
            language: 'es',
            url: cfg.mailChimpURL
        });

        $.ajaxChimp.translations.es = {
            'submit': 'Submitting...',
            0: '<i class="fas fa-check"></i> We have sent you a confirmation email',
            1: '<i class="fas fa-exclamation-triangle"></i> You must enter a valid e-mail address.',
            2: '<i class="fas fa-exclamation-triangle"></i> E-mail address is not valid.',
            3: '<i class="fas fa-exclamation-triangle"></i> E-mail address is not valid.',
            4: '<i class="fas fa-exclamation-triangle"></i> E-mail address is not valid.',
            5: '<i class="fas fa-exclamation-triangle"></i> E-mail address is not valid.'
        }
    };

    (function ssInit() {
        ssPreloader();
        ssInfoToggle();
        ssSlickSlider();
        ssPlaceholder();
        ssFinalCountdown();
        ssAjaxChimp();
    })();

})(jQuery);

async function requestAccess(event) {
    event.preventDefault(); // Предотвращает перезагрузку страницы при отправке формы

    const email = document.getElementById('email').value;

    try {
        const response = await fetch('/request-access', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email }),
        });

        const result = await response.json();
        alert(result.message);
    } catch (error) {
        alert('Произошла ошибка: ' + error.message);
    }
}

// Привязка обработчика событий к форме
document.getElementById('accessForm').addEventListener('submit', requestAccess);
