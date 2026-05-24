<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'

interface Choice {
  id: number
  key: string
  text: string
}

interface Question {
  id: number
  question: string
  choices: Choice[]
}

const errorMessage = ref('')
const name = ref('')
const questions = ref<Question[]>([])
const answers = ref<Record<number, string>>({})
const score = ref<number | null>(null)

onMounted(async () => {
  fetchQuestions()
})

const fetchQuestions = async () => {
  const res = await axios.get(
    'http://localhost:8080/api/questions'
  )

  questions.value = res.data
}

const submitExam = async () => {
  errorMessage.value = ''

  if (!name.value.trim()) {
    errorMessage.value = 'กรุณากรอกชื่อ-สกุล'
    return
  }

  const res = await axios.post(
    'http://localhost:8080/api/submit',
    {
      name: name.value,
      answers: answers.value,
    }
  )

  score.value = res.data.score
}

const resetExam = () => {
  name.value = ''
  answers.value = {}
  score.value = null
  errorMessage.value = ''
}
</script>

<template>
  <div style="padding: 20px">
    <div style="margin-top: 20px">
      <label>
        ชื่อ-สกุล
      </label>

      <input v-model="name"  :disabled="score !== null" required style="
    margin-top: 5px;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 6px;
    background-color: #f5f5f5;
    color: #333;
  " />
      <p v-if="errorMessage" style="color: red">
        {{ errorMessage }}
      </p>
    </div>

    <div v-for="(question, index) in questions" :key="question.id" style="margin-top: 20px">
      <h3>{{ index + 1 }}. {{ question.question }}</h3>

      <div v-for="choice in question.choices" :key="choice.id">
        <label>
          <input type="radio" :name="'question-' + question.id" :value="choice.key" v-model="answers[question.id]" />

          {{ choice.key }} - {{ choice.text }}
        </label>
      </div>
    </div>
    <div style="
    display: flex;
    align-items: left;
    gap: 16px;
    margin-top: 20px;
  ">
      <button v-if="score === null" @click="submitExam" class="action-button">
        ส่งข้อสอบ
      </button>

      <button v-else @click="resetExam" class="action-button">
        สอบใหม่อีกครั้ง
      </button>

      <h5 v-if="score !== null">
        คุณ {{ name }} ได้คะแนน: {{ score }} / {{ questions.length }}
      </h5>
    </div>
  </div>
</template>