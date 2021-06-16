<template>
  <div class="p-2 space-x-2">
    <Btn name="update" @click="updateTaskTree" />
  </div>

  <!-- <div>{{ taskTree }}</div> -->

  <div class="flex flex-row">
    <!-- todo: key to uid -->

    <div class="p-2 space-y-2">
      <Btn name="add task" @click="addTask" />
      <input
        class="block"
        type="text"
        name="name"
        placeholder="task name"
        v-model="newTaskName"
      />

      <TaskCard
        v-for="(item, index) in taskTree"
        :key="index"
        :isActive="enabledTask === index"
        :taskName="item.name"
        :total="item.totalCost"
        :average="item.averageCost"
        @click="selectTask(index)"
      />
    </div>

    <div v-if="enabledTask" class="p-2 space-y-2">
      <Btn name="add subtask" @click="addSubtask" />
      <input
        class="block"
        type="text"
        name="name"
        placeholder="task name"
        v-model="newSubtaskName"
      />

      <SubTaskCard
        v-for="(item, index) in subTaskTree"
        :key="index"
        :isActive="enabledSubTask === index"
        :subtaskName="item.name"
        :total="item.totalCost"
        :average="item.averageCost"
        @click="selectSubTask(index)"
      />
    </div>

    <div class="p-2 space-y-2">
      <div v-if="enabledSubTask != ''">
        <Btn name="add cost" @click="addCost" />
        <input
          class="block"
          type="text"
          placeholder="task name"
          v-model="newCostName"
        />
        <input
          class="block"
          type="number"
          placeholder="task value"
          v-model="newCostValue"
        />
      </div>

      <CostCard
        v-for="(item, index) in costs"
        :key="index"
        :costName="item.name"
        :value="item.value"
      />
    </div>

    <!-- <div v-for="(item, index) in subTaskTree" :key="index">
      {{ item }} - {{ index }}
    </div> -->
  </div>
</template>

<script>
import Btn from './components/Btn.vue'
import CostCard from './components/CostCard.vue'
import SubTaskCard from './components/SubTaskCard.vue'
import TaskCard from './components/TaskCard.vue'
export default {
  components: { TaskCard, SubTaskCard, CostCard, Btn },

  data() {
    return {
      enabledTask: '',
      enabledSubTask: '',
      taskTree: null,
      newTaskName: '',
      newSubtaskName: '',
      newCostName: '',
      newCostValue: 0,
    }
  },

  methods: {
    async addTask() {
      try {
        await fetch('http://localhost:8080/task/add', {
          method: 'POST',
          body: JSON.stringify({ name: this.newTaskName }),
        })
      } catch (error) {
        console.log(error)
        return
      }

      this.updateTaskTree()
    },

    selectTask(num) {
      if (this.enabledTask == num) {
        return
      }

      this.enabledSubTask = ''
      this.enabledTask = num
    },

    async addSubtask() {
      try {
        await fetch('http://localhost:8080/subtask/add', {
          method: 'POST',
          body: JSON.stringify({
            parentName: this.enabledTask,
            subtaskName: this.newSubtaskName,
          }),
        })
      } catch (error) {
        console.log(error)
        return
      }

      this.updateTaskTree()
    },

    selectSubTask(num) {
      this.enabledSubTask = num
    },

    async addCost() {
      try {
        await fetch('http://localhost:8080/cost/add', {
          method: 'POST',
          body: JSON.stringify({
            taskName: this.enabledTask,
            parentName: this.enabledSubTask,
            cost: {
              name: this.newCostName,
              value: Number(this.newCostValue),
            },
          }),
        })
      } catch (error) {
        console.log(error)
        return
      }
      this.updateTaskTree()
    },

    async updateTaskTree() {
      try {
        const res = await fetch('http://localhost:8080/task')
        this.taskTree = await res.json()
      } catch (error) {
        console.log(error)
      }
    },
  },

  computed: {
    subTaskTree() {
      if (!this.enabledTask) {
        return []
      }

      return this.taskTree[this.enabledTask].subtasks
    },

    costs() {
      if (!this.enabledTask || !this.enabledSubTask) {
        return []
      }

      return this.subTaskTree[this.enabledSubTask].costs
    },
  },
}
</script>
