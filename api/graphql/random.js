import findRandom from "../../helper/findRandom";  

const randomResolver = async (root, args) => {
    if (args.type) {
        const type = args.type.split(',');
        const q = { type: type };
        try {
            return await findRandom(q);
        } catch(ex) {
            return ex;
        }
      }
      else if (args.name) {
        const name = args.name.split(',');
        const q = { name: name };
        try {
          return await findRandom(q);
        } catch(ex) {
            return ex;
        }
      }
      else {
          try {
            return await findRandom({});
          } catch(ex) {
              return ex;
          }
      }
};

export default randomResolver;