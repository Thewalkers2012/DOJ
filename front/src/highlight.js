import hljs from 'highlight.js';
import 'highlight.js/styles/vs.css';

const install = (Vue) => {
  Vue.directive('highlight', {
    deep: true,
    bind(el, binding) {
      const targets = el.querySelectorAll('code');

      targets.forEach((target) => {
        if (typeof binding.value === 'string') {
          // if a value is directly assigned to the directive, use this
          // instead of the element content.
          // eslint-disable-next-line no-param-reassign
          target.textContent = binding.value;
        }
        hljs.highlightBlock(target);
      });
    },
    componentUpdated(el, binding) {
      // after an update, re-fill the content and then highlight
      const targets = el.querySelectorAll('code');
      targets.forEach((target) => {
        if (typeof binding.value === 'string') {
          // if a value is directly assigned to the directive, use this
          // instead of the element content.
          // eslint-disable-next-line no-param-reassign
          target.textContent = binding.value;
          hljs.highlightBlock(target);
        }
      });
    },
  });
};

if (window.Vue) {
  // eslint-disable-next-line no-undef
  window.hightlight = highlight;
  Vue.use(install) // eslint-disable-line
}

export default install;
