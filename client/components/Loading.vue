<template>
  <div class="loading-mask">
    <div class="loading-spinner">
      <svg viewBox="25 25 50 50" class="circular">
        <circle cx="50" cy="50" r="20" fill="none" class="path"></circle>
      </svg>
    </div>
  </div>
</template>

<style lang="scss">
.loading-parent--relative {
  position: relative !important;
}

.loading-parent--hidden {
  overflow: hidden !important;
}

.loading-mask {
  position: absolute;
  z-index: 3;
  background-color: hsla(0, 0%, 100%, 0.5);
  margin: 0;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  transition: opacity 0.3s;
}

.loading-mask.is-fullscreen {
  position: fixed;
}

.loading-mask.is-fullscreen .loading-spinner {
  margin-top: -25px;
}

.loading-mask.is-fullscreen .loading-spinner .circular {
  height: 50px;
  width: 50px;
}

.loading-spinner {
  top: 50%;
  margin-top: -21px;
  width: 100%;
  text-align: center;
  position: absolute;
}

.loading-spinner .loading-text {
  color: #409eff;
  margin: 3px 0;
  font-size: 14px;
}

.loading-spinner .circular {
  height: 42px;
  width: 42px;
  animation: loading-rotate 2s linear infinite;
}

.loading-spinner .path {
  animation: loading-dash 1.5s ease-in-out infinite;
  stroke-dasharray: 90, 150;
  stroke-dashoffset: 0;
  stroke-width: 3;
  stroke: #409eff;
  stroke-linecap: round;
}

.loading-spinner i {
  color: #409eff;
}

.loading-fade-enter,
.loading-fade-leave-active {
  opacity: 0;
}

@keyframes loading-rotate {
  to {
    transform: rotate(1turn);
  }
}
@keyframes loading-dash {
  0% {
    stroke-dasharray: 1, 200;
    stroke-dashoffset: 0;
  }
  50% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -40px;
  }
  to {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -120px;
  }
}
.animated {
  -webkit-animation-duration: 300ms;
  animation-duration: 300ms;
}

</style>
