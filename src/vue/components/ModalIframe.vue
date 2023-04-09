<template>
  <div :style="[show ? showModal : null]" class="modal">
    <div class="modal_content">
      <div @click="closeModal()" class="modal_header">
        <h2>Просмотр истории редактирования</h2>
        <button @click="closeModal()" class="close"></button>
      </div>
      <iframe :src="`/iframe?idhtml=${idhtml}`" frameborder="0"></iframe>
    </div>
  </div>
</template>

<script>
export default {
  name: "ModalIframe",
  props: {
    show: Boolean,
    idhtml: Number,
  },
  methods: {
    closeModal() {
      this.$emit("onCloseWindow");
    },
  },
  created() {
    window.onclick = function (event) {
      if (event.target == document.querySelector("modal")) {
        document.querySelector("modal").style.display = "none";
      }
    };
  },
  computed: {
    showModal() {
      return {
        display: "block",
      };
    },
  },
};
</script>

<style>
.modal {
  display: none; /* Hidden by default */
  position: fixed; /* Stay in place */
  z-index: 999; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  padding: 50px;
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0, 0, 0); /* Fallback color */
  background-color: rgba(0, 0, 0, 0.4); /* Black w/ opacity */
}

/* Modal Content/Box */
.modal_content {
  display: flex;
  flex-direction: column;
  background-color: #fefefe;
  margin: 0 auto; /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  width: 90%; /* Could be more or less, depending on screen size */
  height: 100%;
}

.modal_content iframe {
  height: 100%;
  opacity: 1;

  animation-name: opac;
  animation-duration: 1s;
}

@keyframes opac {
  from {
    opacity: 0;
  }
  60% {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.modal_header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #aaa;
  padding-bottom: 10px;
  margin-bottom: 20px;
}

.modal_header h2 {
  margin: 0;
}

/* The Close Button */
.close {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
  position: relative;
}

.close::before {
  content: "\d7";
  font-size: 28px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -54%);
}

.close:hover,
.close:focus {
  color: black;
  background-color: white;
  text-decoration: none;
  cursor: pointer;
}
</style>