<template>
  <div :style="[show ? showModal : null]" class="modal">
    <div class="modal-content">
      <div @click="closeModal()" class="modal_header">
        <h2>Читать вопрос</h2>
        <button @click="closeModal()" class="close"></button>
      </div>
      <div class="modal__content">
        <p>{{ content }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ModalReadQuestion",
  props: {
    content: String,
    show: Boolean,
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

<style scoped>
.modal {
  display: none; /* Hidden by default */
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0, 0, 0); /* Fallback color */
  background-color: rgba(0, 0, 0, 0.4); /* Black w/ opacity */
}

/* Modal Content/Box */
.modal-content {
  display: flex;
  flex-direction: column;
  background-color: #fefefe;
  margin: 15% auto; /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  width: 80%; /* Could be more or less, depending on screen size */
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

.modal__content {
  max-height: 300px;
  overflow-y: auto;
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