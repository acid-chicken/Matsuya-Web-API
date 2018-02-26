import findRandom from "../../helper/findRandom";  

const randomResolver = async (root, args) => {
    if (args.type) {
        const type = args.type.split(',');
        const q = { type: type };
        return await findRandom(q);
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
          return await findRandom({});
      }
};

export default randomResolver;