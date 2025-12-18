class AnimatedLoading extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    const elements = this.dataset.elements; // data-elements attribute
    const width = this.dataset.width; // data-width attribute
    const height = this.dataset.height; // data-height attribute

    for (let i = 0; i < elements; i++) {
      const wrapper = document.createElement('div');
      wrapper.classList.add('loading-wave');
      wrapper.style.width = width;
      wrapper.style.height = height;
      wrapper.style.display = 'inline-block';
      wrapper.style.margin = '10px';
      this.appendChild(wrapper);
    }
  }
}

// Register (inject...) the custom element for use by the DOM HTML
customElements.define('animated-loading', AnimatedLoading);
