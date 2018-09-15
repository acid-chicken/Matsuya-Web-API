import dbcontroller from '../modules/dbcontroller';

const dbSearch = (query) => {
    return new Promise((resolve, reject) => {
        dbcontroller.menu.find(query, function (e, docs) {
            resolve(docs);
            }, err => {
                reject(err);
            });
        });
    };

  export default dbSearch;
