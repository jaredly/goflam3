
var consts = module.exports = {
  name: {
    male: ['James', 'John', 'Peter', 'Sam', 'Kevin', 'Joseph', 'Luke', 'Nephi'],
    female: ['Samantha', 'Jane', 'Judy', 'Anna', 'Maria', 'Lucy', 'Lisa', 'Daphne', 'Pollyanna'],
    last: ['Smith', 'Jorgensen', 'Kaiser', 'Brown', 'Olsen', 'Neuman', 'Frank', 'Schwartz']
  },
  cities: ['Budabest', 'Boston', 'Detroit', 'Paris', 'Athens', 'New Orleans', 'Moscow', 'Berlin', 'San Jose', 'Monta Ray']
}
consts.name.first = consts.name.female.concat(consts.name.male)

