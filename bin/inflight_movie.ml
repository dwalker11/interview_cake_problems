let solution flight_length movie_lengths =
  let result = ref false in
  let table = Hashtbl.create 100 in
  let build_hash = List.iteri (fun i v -> Hashtbl.add table v i) in
  let do_something = List.iteri (fun pos movie_length ->
    let x = flight_length - movie_length in
      match Hashtbl.find table x with
      | v -> if v != pos then result := true
      | exception Not_found -> ()
    )
  in
    build_hash movie_lengths;
    do_something movie_lengths;
    !result

let () =
  let flight_length = 180 in
  let movie_lengths = [90; 90] in
  let result = solution flight_length movie_lengths in
    print_endline (string_of_bool result)