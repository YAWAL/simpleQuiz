<template>
  <div class="quiz">

    <div class="jumbotron">
      <h1>Simple Quiz</h1>
      <p>Code test</p>
    </div>

    <div class="container">
      <div class="row">
        <div class="col-sm"></div>
        <div class="col-4">
          <form v-on:submit="addQuiz">
            <div class="form-group">
              <label><strong>Choose name</strong></label>
              <input type="text" class="form-control" v-model="newQuiz.name" placeholder="Enter your name here" >
            </div>
            <div class="form-group" v-for="(quiz, index) in quizes">
              <label><strong>{{qnum = index+1}} . {{quiz.question}}</strong></label>
              <div class="form-check" v-for="(answer, key) in quiz.answers">
                <input class="form-check-input" type="radio" v-model="newQuiz[index]"  v-bind:value="key">
                <label class="form-check-label">
                  {{key}}: {{answer}}
                </label>
              </div>
            </div>

            <button type="submit" class="btn btn-primary">Submit</button>
          </form>
        </div>
        <div class="col-sm"></div>
      </div>
      <hr>
      <div class="row">
        <div class="col-sm"></div>
        <div class="col-6">
          <span class="text-center" v-for="user in users">
            <h2>Quiz Results</h2>
            <p>Your score <strong>{{correct}}</strong> out of <strong>{{quizes.length}}</strong> questions. Your percentage is <strong>{{perc}}%</strong>.</p>
            <p>You were better than <strong>{{ totalOfAll }}%</strong> of all quizer.</p>
          </span>
        </div>
        <div class="col-sm"></div>
      </div>
    </div>
  </div>

</template>

<script>
  export default {
    name: 'quiz',
    data(){
      return {
        newQuiz: {},
        quizes: [],
        users: [],
        savedUsers:[] ,
        correct:0,
        perc:null,
        oppPerc:null
      }
    },
  methods: {
      addQuiz: function(e){
        this.quizes.forEach((a, index) => {
          if(this.newQuiz[index] === a.correctAnswer) this.correct++;
        });
        this.perc = ((this.correct / this.quizes.length)*100).toFixed(2);
        this.users.push({
          name: this.newQuiz.name,
          score: this.correct
        });
        // console.log('Value1: '+this.savedUsers[0].name+' '+this.savedUsers[0].score);
        // console.log('Value2: '+this.savedUsers[1].name+' '+this.savedUsers[1].score);
        e.preventDefault();
      }
    },
    computed: {
      totalSum: function() {
        let total = 0;
        for(let i = 0; i < this.savedUsers.length; i++){
          total += parseInt(this.savedUsers[i].score);
        }
        return total;
      },
      totalOfAll: function() {
        var sumTotal = (((this.totalSum / 2) / 3)*this.perc).toFixed(2);
        return sumTotal;
      }
    },
    created: function(){
      // this.$http.get('http://localhost:8888/players')
      // .then(function(response){
      //   this.savedUsers = response.data;
      //   console.log(response.data);
      // });

      // quizes json questions and answers https://api.myjson.com/bins/vnryw
      this.$http.get('http://localhost:8888/quiz')
      .then(function(response){
        this.quizes = response.data;
        console.log(response.data);
      });
    }
  }
</script>

<style scoped>
  .jumbotron {
      text-align: center;
      padding: 20px !important;
  }
</style>
