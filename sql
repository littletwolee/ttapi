Create or replace function lrelationship_ike(input_user_id INT, input_other_user_id INT) RETURNS void as
                   $body$
                            Declare
                            ustate VARCHAR;
                            ostate VARCHAR;
                            Begin
                                     SELECT  relationships.state INTO ostate FROM relationships WHERE  relationships.other_user_id = input_user_id AND relationships.user_id = input_other_user_id;
	                                    IF ostate = 'liked' THEN
                                                  UPDATE relationships SET state =  'matched'  WHERE relationships.other_user_id = input_user_id AND relationships.user_id = input_other_user_id;
                                                  SELECT state INTO ustate FROM relationships WHERE  relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id;
                                                  IF ustate IS NULL THEN
                                                                INSERT INTO relationships(user_id, state, type, other_user_id) VALUES(input_user_id, 'matched', 'relationship', input_other_user_id);
																									ELSEIF ustate = 'unliked'  THEN
                                                                UPDATE relationships SET state = 'matched' WHERE relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id ;
                                                   END IF;
                                      ELSE                                      
                                                   SELECT relationships.state INTO ustate FROM relationships WHERE  relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id;
                                                   IF ustate = 'unliked' THEN
																														     UPDATE relationships SET state = 'liked' WHERE relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id ;
                                                   ELSEIF ustate IS NULL THEN
                                                                 INSERT INTO relationships(user_id, state, type, other_user_id) VALUES(input_user_id, 'liked', 'relationship', input_other_user_id);
																									 END IF;
                                       END IF;
                                        
														END;
												$body$
                           LANGUAGE 'plpgsql' VOLATILE;
