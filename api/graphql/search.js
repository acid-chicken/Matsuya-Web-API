import dbSearch from "../../helper/dbSearch";  

const searchResolver = async (root, args) => {
    if (args.type) {
        const type = args.type.split(',');
        const q = { type: type };
        try {
            return await dbSearch(q);
        } catch(ex) {
            return ex;
        }
      }
      else if (args.name) {
        const name = args.name.split(',');
        const q = { name: name };
        try {
          return await dbSearch(q);
        } catch(ex) {
            return ex;
        }
      }
      else {
          try {
            return await dbSearch({});
          } catch(ex) {
              return ex;
          }
      }
};

export default searchResolver;